package validator

import (
	"fmt"
)

type VBool struct {
	Required bool
}

func (vr VBool) CheckValue(v interface{}) error {
	str := fmt.Sprint(v)

	if v == nil || str == "" {
		if vr.Required {
			return &FieldError{FieldRequired}
		} else {
			return nil
		}
	}

	for _, v := range []string{"true", "false", "0", "1"} {
		if str == v {
			return nil
		}
	}

	return &FieldError{FieldWrongType}
}

func (vr VBool) IsRequired() bool {
	return vr.Required
}
