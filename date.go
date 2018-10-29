package validator

import (
	"time"
)

type VDate struct {
	Min      string
	Max      string
	Format   string
	Required bool
}

func (vr VDate) CheckValue(v string) *VFieldResult {
	if len(v) == 0 || v == "null" {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	dt, err := time.Parse(vr.Format, v)

	if err != nil {
		return &VFieldResult{FieldNoDate}
	}

	if vr.Min != "" {
		min, err := time.Parse(vr.Format, vr.Min)

		if err != nil {
			return &VFieldResult{MinNoDate}
		}

		if dt.Before(min) {
			return &VFieldResult{FieldMinVal, vr.Min}
		}
	}

	if vr.Max != "" {
		max, err := time.Parse(vr.Format, vr.Max)

		if err != nil {
			return &VFieldResult{MaxNoDate}
		}

		if dt.After(max) {
			return &VFieldResult{FieldMaxVal, vr.Max}
		}
	}

	return nil
}
