package njson

import (
	"reflect"

	"github.com/tidwall/gjson"
)

func parseStructureType(result gjson.Result, field reflect.Type) (v interface{}) {
	switch field.Kind() {
	case reflect.Slice:
		v = unmarshalSlice(result.Array(), field)
	case reflect.Map:
		v = unmarshalMap(result.Raw, field)
	case reflect.Struct:
		if field.String() == "time.Time" {
			v = result.Time()
		} else {
			v = unmarshalStruct(result.Raw, field)
		}
	default:
		v = nil
	}

	return
}

func parseDataType(result gjson.Result, typ string) (v interface{}) {
	switch typ {
	case reflect.String.String():
		v = result.String()
	case reflect.Int.String():
		v = int(result.Int())
	case reflect.Int8.String():
		v = int8(result.Int())
	case reflect.Int16.String():
		v = int16(result.Int())
	case reflect.Int32.String():
		v = int32(result.Int())
	case reflect.Int64.String():
		v = int64(result.Int())
	case reflect.Float32.String():
		v = float32(result.Float())
	case reflect.Float64.String():
		v = float64(result.Float())
	case reflect.Bool.String():
		v = result.Bool()
	case reflect.Struct.String():
		v = nil
	default:
		v = nil
	}

	return
}
