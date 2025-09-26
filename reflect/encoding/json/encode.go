package json

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Marshal(v any) (string, error) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return "", fmt.Errorf("only support struct")
	}
	typ := reflect.TypeOf(v)

	builder := strings.Builder{}
	builder.WriteString("{")

	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)

		tag := fieldType.Tag.Get("json")
		if tag == "" {
			tag = fieldType.Name
		}

		builder.WriteString(`"` + tag + `":`)

		switch fieldVal.Kind() {
		case reflect.String:
			builder.WriteString(`"` + fieldVal.String() + `"`)
		case reflect.Int:
			builder.WriteString(strconv.FormatInt(fieldVal.Int(), 10))
		default:
			return "", fmt.Errorf("unsupport field type: %s", fieldVal.Kind())
		}

		if i < val.NumField()-1 {
			builder.WriteString(",")
		}
	}

	builder.WriteString("}")
	return builder.String(), nil

}
