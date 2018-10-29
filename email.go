package validator

import (
	"regexp"
)

var regexEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)+$")

type VEmail struct {
	Required bool
}

func (vr VEmail) CheckValue(v string) *VFieldResult {
	if len(v) == 0 || v == "null" {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	if !regexEmail.MatchString(v) {
		return &VFieldResult{FieldNoEmail}
	}

	return nil
}
