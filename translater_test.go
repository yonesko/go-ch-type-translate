package translater

import (
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type Example struct {
		Name int    `json:"name"`
		F4   string `json:"f_4"`
	}

	fmt.Println(translate(reflect.TypeOf(Example{})))
}
