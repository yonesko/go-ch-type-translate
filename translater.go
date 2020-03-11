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
		fieldsStr = append(fieldsStr, fmt.Sprintf("%s %s", fieldName(f), typeTr(f.Type)))
	}
	return strings.Join(fieldsStr, ",\n")
}

func fieldName(f reflect.StructField) string {
	name := f.Tag.Get("json")
	if name == "" {
		name = f.Name
	}
	return name
}

func typeTr(t reflect.Type) string {
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		return fmt.Sprintf("Nested (%s)", translate(t.Elem()))
	}

	switch t.Name() {
	case "int":
		return "Int32"
	default:
		return strings.ToUpper(t.Name()[:1]) + t.Name()[1:]
	}
}
