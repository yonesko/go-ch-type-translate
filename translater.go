package translater

import (
	"fmt"
	"reflect"
	"strings"
)

//Columnize splits slice and array types for Nested type json requirements insert in Clickhouse\
//https://clickhouse.tech/docs/ru/interfaces/formats/#jsoneachrow-nested
func Columnize(t reflect.Type) string {
	s := "type %s struct {\n%s\n}"
	var fieldsStr []string
	for i := 0; i < t.NumField(); i++ {
		fieldsStr = append(fieldsStr, columnize(t.Field(i)))
	}
	return fmt.Sprintf(s, t.Name(), strings.Join(fieldsStr, "\n"))
}

func columnize(f reflect.StructField) string {
	if f.Type.Kind() == reflect.Slice || f.Type.Kind() == reflect.Array {
		t := f.Type.Elem()
		var fieldsStr []string
		for i := 0; i < t.NumField(); i++ {
			fi := t.Field(i)
			fieldsStr = append(fieldsStr,
				fmt.Sprintf("%s%s []%s `%s`", f.Name, fi.Name, fi.Type, fi.Tag))
		}
		return strings.Join(fieldsStr, "\n")
	}
	return fmt.Sprintf("%s %s `%s`", f.Name, f.Type, f.Tag)
}

func TranslateForCreateTable(t reflect.Type) string {
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
		return fmt.Sprintf("Nested (%s)", TranslateForCreateTable(t.Elem()))
	}

	switch t.Name() {
	case "bool":
		return "Uint8"
	case "int":
		return "Int32"
	default:
		return strings.ToUpper(t.Name()[:1]) + t.Name()[1:]
	}
}
