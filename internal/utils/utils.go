package utils

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func LastPart[T protoreflect.Name | protoreflect.FullName](s T) string {
	words := strings.Split(string(s), ".")
	return words[len(words)-1]
}

func PrettyJSON(i []byte) []byte {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, i, "", "    "); err != nil {
		return nil
	}
	return prettyJSON.Bytes()
}
