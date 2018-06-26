package validator

import (
	"regexp"
	"strconv"
)

type VNumeric struct {
	Min      int
	Max      int
	Float    bool
	Negative bool
	Decimals int
	Required bool
}

func (vr VNumeric) CheckValue(v string) *VFieldResult {
	if vr.Required && len(v) == 0 {
		return &VFieldResult{FieldRequired}
	}

	re := regexp.MustCompile("^(-?\\d+)(?:\\.?(\\d*))?$")
	matches := re.FindAllStringSubmatch(v, -1)

	if len(matches) == 0 {
		return &VFieldResult{FieldNoNumeric}
	}

	i, err := strconv.Atoi(matches[0][1]) // integer part

	if err != nil {
		return &VFieldResult{FieldNoNumeric}
	}

	if i < 0 && !vr.Negative {
		return &VFieldResult{FieldIsNegative}
	}

	if matches[0][2] != "" && !vr.Float {
		return &VFieldResult{FieldIsFloat}
	}

	if vr.Float && vr.Decimals > 0 && vr.Decimals < len(matches[0][2]) {
		return &VFieldResult{FieldHasTooManyDecimals, strconv.Itoa(vr.Decimals)}
	}

	if vr.Min == 0 && vr.Max == 0 {
		return nil
	}

	// d, _ := strconv.Atoi(matches[0][2]) // make also decimal part

	if vr.Min != 0 && i < vr.Min {
		return &VFieldResult{FieldMinVal, strconv.Itoa(vr.Min)}
	}

	if vr.Max != 0 && i > vr.Max {
		return &VFieldResult{FieldMaxVal, strconv.Itoa(vr.Max)}
	}

	return nil
}
