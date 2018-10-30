package validator

import (
	"fmt"
	"regexp"
	"strconv"
)

var regexNumeric = regexp.MustCompile("^(-?\\d+)(?:\\.?(\\d*))?$")

type VNumeric struct {
	Min        int
	Max        int
	Minf       float64
	Maxf       float64
	CheckRange bool
	Float      bool
	Negative   bool
	Decimals   int
	Required   bool
}

func (vr VNumeric) CheckValue(v interface{}) *VFieldResult {
	str := fmt.Sprint(v)

	if v == nil || str == "" {
		if vr.Required {
			return &VFieldResult{FieldRequired}
		} else {
			return nil
		}
	}

	matches := regexNumeric.FindAllStringSubmatch(str, -1)

	if len(matches) == 0 {
		return &VFieldResult{FieldNoNumeric}
	}

	i, err := strconv.Atoi(matches[0][1]) // integer part

	if err != nil {
		return &VFieldResult{FieldNoNumeric}
	}

	if vr.Min < 0 || vr.Max < 0 || vr.Minf < 0 || vr.Maxf < 0 {
		vr.Negative = true
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

	if !vr.CheckRange && (!vr.Float && vr.Min == 0 && vr.Max == 0 || vr.Float && vr.Minf == 0 && vr.Maxf == 0) {
		return nil
	}

	// integer comparison
	if !vr.Float {
		if (vr.CheckRange || vr.Min != 0) && i < vr.Min {
			return &VFieldResult{FieldNumMinVal, strconv.Itoa(vr.Min)}
		}

		if (vr.CheckRange || vr.Max != 0) && i > vr.Max {
			return &VFieldResult{FieldNumMaxVal, strconv.Itoa(vr.Max)}
		}
	} else { // float comparison
		f, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return &VFieldResult{FieldNoNumeric}
		}

		if (vr.CheckRange || vr.Minf != 0) && f < vr.Minf {
			return &VFieldResult{FieldNumMinVal, fmt.Sprint(vr.Minf)}
		}

		if (vr.CheckRange || vr.Maxf != 0) && f > vr.Maxf {
			return &VFieldResult{FieldNumMaxVal, fmt.Sprint(vr.Maxf)}
		}
	}

	return nil
}

func (vr VNumeric) IsRequired() bool {
	return vr.Required
}
