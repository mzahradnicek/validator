package validator

import (
	"fmt"
	"time"
)

type VDate struct {
	Min      string
	Max      string
	Format   string
	Required bool
}

func (vr VDate) CheckValue(v interface{}) *VFieldResult {
	str := fmt.Sprint(v)

	if v == nil || str == "" {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	dt, err := time.Parse(vr.Format, str)

	if err != nil {
		return &VFieldResult{FieldNoDate}
	}

	if vr.Min != "" {
		min, err := time.Parse(vr.Format, vr.Min)

		if err != nil {
			return &VFieldResult{MinNoDate}
		}

		if dt.Before(min) {
			return &VFieldResult{FieldDateMinVal, vr.Min}
		}
	}

	if vr.Max != "" {
		max, err := time.Parse(vr.Format, vr.Max)

		if err != nil {
			return &VFieldResult{MaxNoDate}
		}

		if dt.After(max) {
			return &VFieldResult{FieldDateMaxVal, vr.Max}
		}
	}

	return nil
}

func (vr VDate) IsRequired() bool {
	return vr.Required
}
