package validator

import (
	"fmt"
	"regexp"
)

var regexUrl = regexp.MustCompile("^(?:http(s)?:\\/\\/)?[\\w.-]+(?:\\.[\\w\\.-]+)+[\\w\\-\\._~:/?#[\\]@!\\$&'\\(\\)\\*\\+,;=.]+$")

type VUrl struct {
	Required bool
}

func (vr VUrl) CheckValue(v interface{}) *VFieldResult {
	str := fmt.Sprint(v)

	if v == nil || str == "" {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	// make as url.parse

	if !regexUrl.MatchString(str) {
		return &VFieldResult{FieldNoUrl}
	}

	return nil
}

func (vr VUrl) IsRequired() bool {
	return vr.Required
}
