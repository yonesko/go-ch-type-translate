package translater

import (
	"fmt"
	"reflect"
	"testing"
)

type ElemExample struct {
	CompanyInn string `json:"companyInn"`
	CompanyKPP string `json:"companyKPP"`
}
type Example struct {
	Name  int           `json:"name"`
	F4    string        `json:"f_4"`
	Elems []ElemExample `json:"elems"`
}

func TestTranslateForCreateTable(t *testing.T) {
	fmt.Println(TranslateForCreateTable(reflect.TypeOf(Example{})))
}

func TestColumnize(t *testing.T) {
	fmt.Println(Columnize(reflect.TypeOf(Example{})))
}
