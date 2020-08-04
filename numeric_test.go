package validator

import (
	"reflect"
	"testing"
)

func TestVNumeric(t *testing.T) {
	tables := []struct {
		validator VNumeric
		v         interface{}
		res       *FieldError
	}{
		{VNumeric{}, nil, nil},

		// Required
		{VNumeric{Required: true}, nil, &FieldError{FieldRequired}},
		{VNumeric{Required: true}, "", &FieldError{FieldRequired}},

		// Field type
		{VNumeric{}, 1234, nil},
		{VNumeric{}, "1234", nil},
		{VNumeric{}, "asdf", &FieldError{FieldNoNumeric}},
		{VNumeric{}, "1234.234", &FieldError{FieldIsFloat}},
		{VNumeric{Float: true}, "1234.234", nil},

		// Negative
		{VNumeric{}, "-1234", &FieldError{FieldIsNegative}},
		{VNumeric{Negative: true}, "-1234", nil},
		{VNumeric{Float: true}, "-1234.2343", &FieldError{FieldIsNegative}},
		{VNumeric{Float: true, Negative: true}, "-1234.234234", nil},

		// Decimals
		{VNumeric{Float: true, Negative: true, Decimals: 3}, "-1234.23433234", &FieldError{FieldHasTooManyDecimals, "3"}},
		{VNumeric{Float: true, Negative: true, Decimals: 3}, "-1234.234", nil},

		// Min integer
		{VNumeric{Min: 10, Negative: true}, -1234, &FieldError{FieldNumMinVal, "10"}},
		{VNumeric{Min: 10}, 4, &FieldError{FieldNumMinVal, "10"}},
		{VNumeric{Min: -20, Negative: true}, -40, &FieldError{FieldNumMinVal, "-20"}},
		{VNumeric{Min: 10}, 12, nil},
		{VNumeric{Min: -10, Negative: true}, -5, nil},

		// Max integer
		{VNumeric{Max: -10}, 1234, &FieldError{FieldNumMaxVal, "-10"}},
		{VNumeric{Max: 4}, 10, &FieldError{FieldNumMaxVal, "4"}},
		{VNumeric{Max: -40, Negative: true}, -20, &FieldError{FieldNumMaxVal, "-40"}},
		{VNumeric{Max: 12}, 10, nil},
		{VNumeric{Max: 12}, 12, nil},
		{VNumeric{Max: -5, Negative: true}, -10, nil},

		// Min float
		{VNumeric{Minf: 10.234, Float: true, Negative: true}, -1234.234234, &FieldError{FieldNumMinVal, "10.234"}},
		{VNumeric{Minf: 10.9558, Float: true}, 4, &FieldError{FieldNumMinVal, "10.9558"}},
		{VNumeric{Minf: -20.234555, Float: true, Negative: true}, -40.23423, &FieldError{FieldNumMinVal, "-20.234555"}},
		{VNumeric{Minf: 10.234555, Float: true}, 12.23423, nil},
		{VNumeric{Minf: -10.98747, Float: true, Negative: true}, -5.854566, nil},

		// Max float
		{VNumeric{Maxf: -10.2344, Float: true}, 1234.23423, &FieldError{FieldNumMaxVal, "-10.2344"}},
		{VNumeric{Maxf: 4.35544, Float: true}, "10.32423", &FieldError{FieldNumMaxVal, "4.35544"}},
		{VNumeric{Maxf: -40.34523424, Float: true, Negative: true}, -20.232334, &FieldError{FieldNumMaxVal, "-40.34523424"}},
		{VNumeric{Maxf: 12.2345564, Float: true}, 10.2383848, nil},
		{VNumeric{Maxf: 12.88865, Float: true}, 12.3658435, nil},
		{VNumeric{Maxf: -5.878474, Float: true, Negative: true}, -10.5684684, nil},

		// CheckRange
		{VNumeric{Min: 0, Max: 25, CheckRange: true, Negative: true}, -1234, &FieldError{FieldNumMinVal, "0"}},
		{VNumeric{Min: -25, Max: 0, CheckRange: true}, 1234, &FieldError{FieldNumMaxVal, "0"}},

		{VNumeric{Minf: 0, Maxf: 25.23443, Float: true, CheckRange: true, Negative: true}, -1234.234234, &FieldError{FieldNumMinVal, "0"}},
		{VNumeric{Minf: -24.345, Maxf: 0, Float: true, CheckRange: true}, 12.234234, &FieldError{FieldNumMaxVal, "0"}},

		// write some more cases
	}

	for _, table := range tables {
		var res *FieldError
		eres := table.validator.CheckValue(table.v)
		if eres != nil {
			res = eres.(*FieldError)
		}

		if (table.res != res && (table.res == nil || res == nil)) || (table.res != nil && res != nil && !reflect.DeepEqual(*res, *table.res)) {
			t.Errorf("Text validator %+v for \"%v\" got:  \"%v\", want: \"%v\".", table.validator, table.v, res, table.res)
		}
	}
}
