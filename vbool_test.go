package validator

import (
	"reflect"
	"testing"
)

func TestVBool(t *testing.T) {
	tables := []struct {
		validator VBool
		v         interface{}
		res       *FieldError
	}{
		{VBool{}, nil, nil},

		// Required
		{VBool{Required: true}, nil, &FieldError{FieldRequired}},
		{VBool{Required: true}, "", &FieldError{FieldRequired}},

		// URL Format
		{VBool{Required: true}, true, nil},
		{VBool{Required: true}, false, nil},
		{VBool{Required: true}, 1, nil},
		{VBool{Required: true}, 0, nil},
		{VBool{Required: true}, "otherval", &FieldError{FieldWrongType}},
		{VBool{Required: true}, 10, &FieldError{FieldWrongType}},
	}

	for _, table := range tables {
		var res *FieldError
		eres := table.validator.CheckValue(table.v)
		if eres != nil {
			res = eres.(*FieldError)
		}

		if (table.res != res && (table.res == nil || res == nil)) || (table.res != nil && res != nil && !reflect.DeepEqual(*res, *table.res)) {
			t.Errorf("Bool validator %+v for \"%v\" got:  \"%v\", want: \"%v\".", table.validator, table.v, res, table.res)
		}
	}
}
