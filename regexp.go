package validator

import (
	"fmt"
	"regexp"
)

type VRegexp struct {
	Pattern  string
	Required bool
	regexp   *regexp.Regexp
}

func (vr VRegexp) CheckValue(v interface{}) *VFieldResult {
	str := fmt.Sprint(v)

	if v == nil || str == "" {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	// compile regular expression
	if vr.regexp == nil {
		re, err := regexp.Compile(vr.Pattern)

		if err != nil {
			return &VFieldResult{FieldBadRegexp}
		}

		vr.regexp = re
	}

	if !vr.regexp.MatchString(str) {
		return &VFieldResult{FieldNoMatch}
	}

	return nil
}

func (vr VRegexp) IsRequired() bool {
	return vr.Required
}
