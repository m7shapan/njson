package njson

import "reflect"

func validTag(filed reflect.StructField) bool {
	return !(filed.Tag.Get(tag) == "" ||
		filed.Tag.Get(tag) == "-")
}

func isStructureType(typ string) (ok bool) {
	switch typ {
	case reflect.Slice.String():
		ok = true
	case reflect.Map.String():
		ok = true
	case reflect.Struct.String():
		ok = true
	default:
		ok = false
	}

	return
}
