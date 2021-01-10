package njson

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/tidwall/gjson"
)

const tag string = "njson"

// Unmarshal used to unmarshal nested json using "njson" tag
func Unmarshal(data []byte, v interface{}) (err error) {
	// catch code panic and return error message
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		}
	}()

	elem := reflect.ValueOf(v).Elem()
	typeOfT := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)

		if !(validTag(typeOfT.Field(i)) && field.CanSet()) {
			continue
		}

		// get field value by tag
		result := gjson.GetBytes(data, typeOfT.Field(i).Tag.Get(tag))

		var value interface{}
		if isStructureType(field.Kind().String()) {
			value = parseStructureType(result, field.Type())
		} else {
			// set field value depend on it's data type
			value = parseDataType(result, field.Kind().String())
		}

		if v != nil {
			elem.Field(i).Set(reflect.ValueOf(value))
		}
	}

	return nil
}

func unmarshalSlice(results []gjson.Result, field reflect.Type) interface{} {
	newSlice := reflect.MakeSlice(field, 0, 0)

	for i := 0; i < len(results); i++ {

		r := results[i]
		k := reflect.TypeOf(r)

		var value interface{}
		if isStructureType(k.String()) {
			value = parseStructureType(r, field.Elem())
		} else {
			value = parseDataType(r, field.Elem().String())
		}

		if value != nil {
			newSlice = reflect.Append(newSlice, reflect.ValueOf(value))
		}
	}

	return newSlice.Interface()
}

func unmarshalMap(raw string, field reflect.Type) interface{} {
	m := reflect.New(reflect.MapOf(field.Key(), field.Elem())).Interface()

	err := json.Unmarshal([]byte(raw), m)
	if err != nil {
		panic(err)
	}

	return reflect.Indirect(reflect.ValueOf(m)).Interface()
}

func unmarshalStruct(raw string, field reflect.Type) interface{} {
	v := reflect.New(field).Interface()

	err := Unmarshal([]byte(raw), v)
	if err != nil {
		panic(err)
	}

	return reflect.Indirect(reflect.ValueOf(v)).Interface()
}
