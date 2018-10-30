package validator

import (
	"fmt"
	"regexp"
)

var regexEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$")

type VEmail struct {
	Required bool
}

func (vr VEmail) CheckValue(v interface{}) *VFieldResult {
	str := fmt.Sprint(v)

	if v == nil || str == "" {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	if !regexEmail.MatchString(str) {
		return &VFieldResult{FieldNoEmail}
	}

	return nil
}

func (vr VEmail) IsRequired() bool {
	return vr.Required
}
