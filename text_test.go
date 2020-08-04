package validator

import (
	"reflect"
	"testing"
)

func TestVText(t *testing.T) {
	tables := []struct {
		validator VText
		v         interface{}
		res       *FieldError
	}{
		{VText{}, nil, nil},

		// Required
		{VText{Required: true}, nil, &FieldError{FieldRequired}},
		{VText{Required: true}, "", &FieldError{FieldRequired}},

		// Min
		{VText{Min: 15}, "Short text", &FieldError{FieldTextMinVal, "15"}},
		{VText{Min: 15}, "Very very long text", nil},

		// Max
		{VText{Max: 15}, "Very very long text", &FieldError{FieldTextMaxVal, "15"}},
		{VText{Max: 15}, "Short text", nil},
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
