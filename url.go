package validator

import (
	"regexp"
)

var regexUrl = regexp.MustCompile("^(?:http(s)?:\\/\\/)?[\\w.-]+(?:\\.[\\w\\.-]+)+[\\w\\-\\._~:/?#[\\]@!\\$&'\\(\\)\\*\\+,;=.]+$")

type VUrl struct {
	Required bool
}

func (vr VUrl) CheckValue(v string) *VFieldResult {
	if len(v) == 0 || v == "null" {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	// make as url.parse

	if !regexUrl.MatchString(v) {
		return &VFieldResult{FieldNoUrl}
	}

	return nil
}
