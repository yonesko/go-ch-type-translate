package translater

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type ElemExample struct {
		CompanyInn string `json:"companyInn"`
		CompanyKPP string `json:"companyKPP"`
	}
	type Example struct {
		Name  int           `json:"name"`
		F4    string        `json:"f_4"`
		Elems []ElemExample `json:"elems"`
	}

	fmt.Println(TranslateForCreateTable(reflect.TypeOf(Example{})))
}
