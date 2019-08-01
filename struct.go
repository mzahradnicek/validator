package validator

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type VStruct struct {
	Rules    VRules
	Required bool
}

func (vr VStruct) CheckValue(v interface{}) *VFieldResult {
	// validate
	str := fmt.Sprint(v)

	if v == nil || str == "" {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	return nil
}

func (vr VStruct) IsRequired() bool {
	return vr.Required
}
