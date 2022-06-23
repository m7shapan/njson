package njson

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/tidwall/gjson"
)

const (
	njsonTag = "njson"
	jsonTag  = "json"
)

var jsonNumberType = reflect.TypeOf(json.Number(""))

// Unmarshal used to unmarshal nested json using "njson" tag
func Unmarshal(data []byte, v interface{}) (err error) {
	if !gjson.ValidBytes(data) {
		return fmt.Errorf("invalid json: %v", string(data))
	}

	// catch code panic and return error message
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = fmt.Errorf("unknown panic: %v", r)
			}
		}
	}()

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("can't unmarshal to invalid type %v", reflect.TypeOf(v))
	}
	elem := rv.Elem()
	typeOfT := elem.Type()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)

		// Check that the tag is either "json" or "njson", and can be set
		if (!validTag(typeOfT.Field(i), njsonTag) && !validTag(typeOfT.Field(i), jsonTag)) || !field.CanSet() {
			continue
		}

		// Assume "njson" by default, but change to "json" if tag matches
		fieldName := typeOfT.Field(i).Tag.Get(njsonTag)
		if validTag(typeOfT.Field(i), jsonTag) {
			fieldName = typeOfT.Field(i).Tag.Get(jsonTag)

			// Only support true "json" tags:
			// if a tag is nested, it must use the "njson" tag
			if len(strings.Split(fieldName, ".")) > 1 {
				return fmt.Errorf("invalid json tag: %s", fieldName)
			}
		}

		// get field value by tag
		result := gjson.GetBytes(data, fieldName)

		// if field type json.Number
		if v != nil && field.Kind() == reflect.String && field.Type() == jsonNumberType {
			elem.Field(i).SetString(result.String())
			continue
		}

		var value interface{}
		if isStructureType(field.Kind().String()) {
			value = parseStructureType(result, field.Type())
		} else {
			// set field value depend on it's data type
			value = parseDataType(result, field.Type().String())
		}

		// maybe it is a custom type, use json.unmarshal
		if value == nil {
			value = unmarshalGeneric(result.Raw, field)
		}

		if v != nil {
			elem.Field(i).Set(reflect.ValueOf(value))
		}
	}

	return
}

func unmarshalSlice(results []gjson.Result, field reflect.Type) interface{} {
	newSlice := reflect.MakeSlice(field, 0, 0)

	for i := 0; i < len(results); i++ {

		var value interface{}
		if isStructureType(field.Elem().Kind().String()) {
			value = parseStructureType(results[i], field.Elem())
		} else {
			// set field value depend on it's data type
			value = parseDataType(results[i], field.Elem().String())
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

func unmarshalGeneric(raw string, field reflect.Value) interface{} {
	err := json.Unmarshal([]byte(raw), field.Addr().Interface())
	if err != nil {
		panic(err)
	}

	return reflect.Indirect(field).Interface()
}
