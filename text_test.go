package validator

import (
	"reflect"
	"testing"
)

func TestVText(t *testing.T) {
	tables := []struct {
		validator VText
		v         interface{}
		res       *VFieldResult
	}{
		{VText{}, nil, nil},

		// Required
		{VText{Required: true}, nil, &VFieldResult{FieldRequired}},
		{VText{Required: true}, "", &VFieldResult{FieldRequired}},

		// Min
		{VText{Min: 15}, "Short text", &VFieldResult{FieldTextMinVal, "15"}},
		{VText{Min: 15}, "Very very long text", nil},

		// Max
		{VText{Max: 15}, "Very very long text", &VFieldResult{FieldTextMaxVal, "15"}},
		{VText{Max: 15}, "Short text", nil},
	}

	for _, table := range tables {
		if res := table.validator.CheckValue(table.v); (table.res != res && (table.res == nil || res == nil)) || (table.res != nil && res != nil && !reflect.DeepEqual(*res, *table.res)) {
			t.Errorf("Text validator %+v for \"%v\" got:  \"%v\", want: \"%v\".", table.validator, table.v, res, table.res)
		}
	}
}
