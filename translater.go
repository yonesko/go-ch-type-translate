package translater

import (
	"fmt"
	"reflect"
	"strings"
)

func translate(t reflect.Type) string {
	var fieldsStr []string
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fieldsStr = append(fieldsStr, fmt.Sprintf("%s %s", f.Tag.Get("json"), typeTr(f.Type)))
	}
	return strings.Join(fieldsStr, ",\n")
}

func typeTr(t reflect.Type) string {
	switch t.Name() {
	case "int":
		return "Int32"
	default:
		return strings.ToUpper(t.Name()[:1]) + t.Name()[1:]
	}
}