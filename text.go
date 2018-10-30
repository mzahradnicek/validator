package validator

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type VText struct {
	Min       int
	Max       int
	Required  bool
	RawLength bool
}

func (vr VText) CheckValue(v interface{}) *VFieldResult {
	str := fmt.Sprint(v)

	if v == nil || str == "" {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	if vr.Min > 0 && ((vr.RawLength && len(str) < vr.Min) || (!vr.RawLength && utf8.RuneCountInString(str) < vr.Min)) {
		return &VFieldResult{FieldTextMinVal, strconv.Itoa(vr.Min)}
	}

	if vr.Max > 0 && ((vr.RawLength && len(str) > vr.Max) || (!vr.RawLength && utf8.RuneCountInString(str) > vr.Max)) {
		return &VFieldResult{FieldTextMaxVal, strconv.Itoa(vr.Max)}
	}

	return nil
}

func (vr VText) IsRequired() bool {
	return vr.Required
}
