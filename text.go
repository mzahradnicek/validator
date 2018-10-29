package validator

import "strconv"
import "unicode/utf8"

type VText struct {
	Min       int
	Max       int
	Multiple  bool
	Required  bool
	RawLength bool
}

func (vr VText) CheckValue(v string) *VFieldResult {
	if len(v) == 0 || v == "null" {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	if vr.Min > 0 && ((vr.RawLength && len(v) < vr.Min) || (!vr.RawLength && utf8.RuneCountInString(v) < vr.Min)) {
		return &VFieldResult{FieldMinVal, strconv.Itoa(vr.Min)}
	}

	if vr.Max > 0 && ((vr.RawLength && len(v) > vr.Max) || (!vr.RawLength && utf8.RuneCountInString(v) > vr.Max)) {
		return &VFieldResult{FieldMaxVal, strconv.Itoa(vr.Max)}
	}

	return nil
}
