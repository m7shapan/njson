package njson

import (
	"reflect"
	"strings"
)

func validTag(filed reflect.StructField, tag string) bool {
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

	if strings.Contains(typ, "[]") {
		ok = true
	}

	return
}
