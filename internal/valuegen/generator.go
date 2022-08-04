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
