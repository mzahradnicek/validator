package validator

type VStruct struct {
	Rules    VRules
	Required bool
}

func (vr VStruct) CheckValue(v interface{}) error {
	if v == nil && vr.Required {
		return &FieldError{FieldRequired}
	}

	res := vr.Rules.Validate(v)

	return res
}

func (vr VStruct) IsRequired() bool {
	return vr.Required
}
