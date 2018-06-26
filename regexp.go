package validator

import (
	"regexp"
)

type VRegexp struct {
	Pattern  string
	Required bool
	regexp   *regexp.Regexp
}

func (vr VRegexp) CheckValue(v string) *VFieldResult {

	// compile regular expression
	if vr.regexp == nil {
		re, err := regexp.Compile(vr.Pattern)

		if err != nil {
			return &VFieldResult{FieldBadRegexp}
		}

		vr.regexp = re
	}

	if len(v) == 0 {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	if !vr.regexp.MatchString(v) {
		return &VFieldResult{FieldNoMatch}
	}

	return nil
}
