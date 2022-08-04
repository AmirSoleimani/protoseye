package protoparser

import (
	"encoding/json"

	"github.com/AmirSoleimani/protoseye/internal/utils"
)

type Output struct {
	Request  map[string]any
	Response map[string]any
}

func (o *Output) MarshalRequest() []byte {
	return o.marshalPretty(o.Request)
}

func (o *Output) MarshalResponse() []byte {
	return o.marshalPretty(o.Response)
}

func (o *Output) marshalPretty(i any) []byte {
	b, _ := json.Marshal(i)
	return utils.PrettyJSON(b)
}
