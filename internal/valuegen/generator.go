package valuegen

import (
	"fmt"
)

// RandomValue generates a random value for the given type.
// MARK: We might want to use `reflect.Kind` types instead of string.
func RandomValue(fieldKind string, isList bool) (any, error) {
	var output any

	switch fieldKind {
	case "bool", "google.protobuf.BoolValue":
		output = true
	case "int32", "sint32", "google.protobuf.Int32Value":
		output = int32(32)
	case "uint32", "google.protobuf.UInt32Value":
		output = uint32(132)
	case "int64", "sint64", "google.protobuf.Int64Value":
		output = int64(64)
	case "uint64", "google.protobuf.UInt64Value":
		output = uint64(64)
	case "sfixed32", "fixed32":
		output = int32(323)
	case "float", "google.protobuf.FloatValue":
		output = float32(25.6)
	case "sfixed64", "fixed64":
		output = int64(644)
	case "double", "google.protobuf.DoubleValue":
		output = 37.6
	case "string", "google.protobuf.StringValue":
		output = "mystring"
	case "bytes", "google.protobuf.BytesValue":
		output = []byte("mybytes")
	case "google.protobuf.Struct":
		output = map[string]any{"key1": "value1", "key2": "value2"}
	case "google.protobuf.Any":
		output = map[string]any{"@type": "string", "value": "yyy"}
	case "google.protobuf.TimeOfDay":
		output = map[string]any{
			"hours":   22,
			"minutes": 15,
			"seconds": 23,
			"nanos":   222,
		}
	case "google.protobuf.Color":
		output = map[string]any{
			"red":   0.2,
			"green": 0.3,
			"blue":  0.4,
			"alpha": map[string]any{
				"value": 1.0,
			},
		}
	case "google.protobuf.Date":
		output = map[string]any{
			"year":  1993,
			"month": 11,
			"day":   27,
		}
	case "google.protobuf.LatLng":
		output = map[string]any{
			"latitude":  85.23,
			"longitude": -23.44,
		}
	case "google.protobuf.Money":
		output = map[string]any{
			"currency_code": "EUR",
			"units":         1,
			"nanos":         750000000,
		}
	case "google.protobuf.Empty":
		output = nil
	default:
		return nil, fmt.Errorf("unsupported field kind: %s", fieldKind)
	}

	if isList {
		output = []any{output}
	}

	return output, nil
}
