package validator

import (
	"reflect"
	"testing"
)

func TestVDate(t *testing.T) {
	tables := []struct {
		validator VDate
		v         interface{}
		res       *FieldError
	}{
		{VDate{}, nil, nil},

		// Required
		{VDate{Required: true}, nil, &FieldError{FieldRequired}},
		{VDate{Required: true}, "", &FieldError{FieldRequired}},

		// write some cases
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
