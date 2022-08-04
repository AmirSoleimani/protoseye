package protoparser

import (
	"fmt"
	"sync"

	"github.com/AmirSoleimani/protoseye/internal/utils"
	"github.com/AmirSoleimani/protoseye/internal/valuegen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Parser struct {
	msgRegistry sync.Map
}

func New() (*Parser, error) {
	return &Parser{}, nil
}

func (p *Parser) RegisterService(service *protogen.Service) (map[string]Output, error) {
	var (
		err  error
		rpcs = map[string]Output{}
	)

	for _, method := range service.Methods {
		rpc := Output{}

		if rpc.Request, err = p.procMessage(method.Input); err != nil {
			return nil, err
		}

		if rpc.Response, err = p.procMessage(method.Output); err != nil {
			return nil, err
		}

		rpcs[string(method.Desc.FullName())] = rpc
	}

	return rpcs, nil
}

func (p *Parser) procMessage(message *protogen.Message) (map[string]any, error) {
	identicalName := message.Desc.FullName()

	// To avoid infinite recursion, we check if the schema is already generated.
	if v, ok := p.msgRegistry.Load(identicalName); ok {
		return v.(map[string]any), nil
	}

	jsonPayload := map[string]any{}
	p.msgRegistry.Store(identicalName, jsonPayload)

	for _, field := range message.Fields {
		v, _ := p.procField(field)
		jsonSC := utils.ToSnakeCase(field.Desc.JSONName())
		jsonPayload[jsonSC] = v
	}

	return jsonPayload, nil
}

func (p *Parser) procField(field *protogen.Field) (any, error) {
	var value any
	var err error

	switch field.Desc.Kind() {
	case protoreflect.EnumKind:
		value, err = p.procFieldEnum(field)
		if err != nil {
			return nil, fmt.Errorf("field %s: couldn't process the enum %s", field.Desc.FullName(), field.Desc.Kind())
		}

	case protoreflect.MessageKind:
		value, err = p.procFieldMessage(field)

	default:
		value, err = valuegen.RandomValue(
			field.Desc.Kind().String(),
			field.Desc.IsList(),
		)
	}

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (p *Parser) procFieldEnum(field *protogen.Field) (any, error) {
	var enumSample any

	if len(field.Enum.Values) > 0 {
		enumSample = utils.LastPart(field.Enum.Values[0].Desc.FullName())
	}

	return enumSample, nil
}

func (p *Parser) procFieldMessage(field *protogen.Field) (any, error) {
	if t, err := valuegen.RandomValue(
		string(field.Message.Desc.FullName()),
		field.Desc.IsList(),
	); err == nil {
		return t, nil
	}

	if field.Desc.IsMap() {
		return p.procFieldMap(field)
	}

	if v, err := p.procMessage(field.Message); err == nil {
		return v, nil
	}

	if field.Desc.IsList() {
		return []any{"a"}, nil
	}

	return nil, nil
}

func (p *Parser) procFieldMap(field *protogen.Field) (any, error) {
	k := field.Desc.MapKey().Kind()

	// JSON format does not support non-string keys.
	// We return an empty map instead.
	if k != protoreflect.StringKind {
		return map[string]any{}, nil
	}

	var vs any
	val := field.Desc.MapValue()
	for _, ff := range field.Message.Fields {
		if ff.Desc == val {
			var err error
			if vs, err = p.procField(ff); err != nil {
				return nil, err
			}
			break
		}
	}

	if vs == nil {
		return nil, fmt.Errorf("field %s: map value type %s not found", field.Desc.FullName(), val.FullName())
	}

	// TODO: improve it
	return map[string]any{
		"mycustomkey": vs,
	}, nil
}
