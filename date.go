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

func (vr VDate) CheckValue(v interface{}) error {
	str := fmt.Sprint(v)

	if v == nil || str == "" {
		if vr.Required {
			return &FieldError{FieldRequired}
		} else {
			return nil
		}
	}

	dt, err := time.Parse(vr.Format, str)

	if err != nil {
		return &FieldError{FieldNoDate}
	}

	if vr.Min != "" {
		min, err := time.Parse(vr.Format, vr.Min)

		if err != nil {
			return &FieldError{MinNoDate}
		}

		if dt.Before(min) {
			return &FieldError{FieldDateMinVal, vr.Min}
		}
	}

	if vr.Max != "" {
		max, err := time.Parse(vr.Format, vr.Max)

		if err != nil {
			return &FieldError{MaxNoDate}
		}

		if dt.After(max) {
			return &FieldError{FieldDateMaxVal, vr.Max}
		}
	}

	return nil
}

func (vr VDate) IsRequired() bool {
	return vr.Required
}
