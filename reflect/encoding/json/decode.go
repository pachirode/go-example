package json

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Unmarshal(data string, v interface{}) error {
	parsedData, err := parseJson(data)
	if err != nil {
		println(err)
		return err
	}

	val := reflect.ValueOf(v).Elem()
	if val.Kind() != reflect.Struct {
		return fmt.Errorf("only support struct")
	}

	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		fieldType := typ.Field(i)

		tag := fieldType.Tag.Get("json")
		if tag == "" {
			tag = fieldType.Name
		}

		if value, ok := parsedData[tag]; ok {
			switch fieldValue.Kind() {
			case reflect.String:
				fieldValue.SetString(value)
			case reflect.Int:
				intValue, err := strconv.Atoi(value)
				if err != nil {
					return err
				}
				fieldValue.SetInt(int64(intValue))
			default:
				return fmt.Errorf("unsupported field type: %s", fieldValue.Kind())
			}
		}
	}

	return nil
}

func parseJson(data string) (map[string]string, error) {
	result := make(map[string]string)

	data = strings.TrimSpace(data)
	if len(data) < 2 || data[0] != '{' || data[len(data)-1] != '}' {
		return nil, fmt.Errorf("invaild JSON")
	}

	data = data[1 : len(data)-1]
	for _, ele := range strings.Split(data, ",") {
		kv := strings.Split(ele, ":")
		if len(kv) != 2 {
			return nil, fmt.Errorf("invaild JSON")
		}

		k := strings.Trim(kv[0], `"`)
		v := strings.Trim(kv[1], `"`)

		result[k] = v
	}

	return result, nil
}
