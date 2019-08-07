package validator

import "fmt"

type VStruct struct {
	Rules    VRules
	Required bool
}

func (vr VStruct) CheckValue(v interface{}) error {
	if v == nil && vr.Required {
		return &FieldError{FieldRequired}
	}

	fmt.Printf("Validate struct\n")

	res := vr.Rules.Validate(v)

	fmt.Printf("%#v\n", res)

	return res
}

func (vr VStruct) IsRequired() bool {
	return vr.Required
}
