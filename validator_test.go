package validator

import (
	"encoding/json"
	"fmt"
	// "reflect"
	"testing"
)

func TestValidate(t *testing.T) {

	rules := VRules{
		"name":   VText{Required: true},
		"age":    VNumeric{Required: true, Min: 18},
		"street": VText{Required: true},
	}

	jsonText := "{\"name\": \"John Doe\", \"age\": 15, \"street\": \"\"}"
	data := map[string]interface{}{}

	err := json.Unmarshal([]byte(jsonText), &data)
	if err != nil {
		fmt.Println("JSON Error: ", err)
	}

	// wants := VResults{}

	res := Validate(data, rules)
	fmt.Printf("res: %#v\n\n", res)

	/*
		for k, v := range res {
			fmt.Printf("Key: %v, Value: %+v\n", k, v[0])
		}
	*/

	/*
		for _, table := range tables {
			if res := table.validator.CheckValue(table.v); (res == nil && table.res != nil) || (res != nil && !reflect.DeepEqual(*res, *table.res)) {
				t.Errorf("Text validator %+v for \"%v\" got:  \"%v\", want: \"%v\".", table.validator, table.v, res, table.res)
			}
		}
	*/
}
