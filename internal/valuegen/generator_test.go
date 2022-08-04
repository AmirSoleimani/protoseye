package valuegen_test

import (
	"reflect"
	"testing"

	"github.com/AmirSoleimani/protoseye/internal/valuegen"
)

const nilK = "nil"

func TestRandomValue(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		fieldKind string
		isList    bool
		wantKind  string
	}{
		{
			fieldKind: "string",
		},
		{
			fieldKind: "int32",
		},
		{
			fieldKind: "bool",
		},
		{
			fieldKind: "uint64",
		},
		{
			fieldKind: "uint32",
		},
		{
			fieldKind: "fixed32",
			wantKind:  "int32",
		},
		{
			fieldKind: "fixed64",
			wantKind:  "int64",
		},
		{
			fieldKind: "double",
			wantKind:  "float64",
		},
		{
			fieldKind: "google.protobuf.Int64Value",
			wantKind:  "int64",
		},
		{
			fieldKind: "int32",
			isList:    true,
		},
		{
			fieldKind: "string",
			isList:    true,
		},
		{
			fieldKind: "bytes",
			wantKind:  "[]uint8",
		},
		{
			fieldKind: "google.protobuf.FloatValue",
			wantKind:  "float32",
		},
		{
			fieldKind: "google.protobuf.Struct",
			wantKind:  "map[string]interface {}",
		},
		{
			fieldKind: "google.protobuf.Any",
			wantKind:  "map[string]interface {}",
		},
		{
			fieldKind: "google.protobuf.Empty",
			wantKind:  nilK,
		},
	}

	for _, tc := range testCases {
		if tc.wantKind == "" {
			tc.wantKind = tc.fieldKind
		}

		v, err := valuegen.RandomValue(tc.fieldKind, tc.isList)
		if err != nil {
			t.Errorf("RandomValue(%s, %t) returned error: %v", tc.fieldKind, tc.isList, err)
			continue
		}

		if v == nil && tc.wantKind == nilK {
			continue
		}

		if tc.isList {
			sliceT := reflect.TypeOf(v)
			if sliceT.Kind() != reflect.Slice {
				t.Errorf("RandomValue(%s, %t) returned %v, want slice", tc.fieldKind, tc.isList, v)
			}

			sliceV := reflect.ValueOf(v)
			v = sliceV.Index(0).Interface()
		}

		if reflect.TypeOf(v).String() != tc.wantKind {
			t.Errorf(
				"RandomValue(%s, %t) returned type %v, want %v",
				tc.fieldKind,
				tc.isList,
				reflect.TypeOf(v).String(),
				tc.wantKind,
			)
		}
	}
}
