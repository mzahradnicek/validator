package validator

import (
	"reflect"
	"testing"
)

func TestVEmail(t *testing.T) {
	tables := []struct {
		validator VEmail
		v         interface{}
		res       *FieldError
	}{
		{VEmail{}, nil, nil},

		// Required
		{VEmail{Required: true}, nil, &FieldError{FieldRequired}},
		{VEmail{Required: true}, "", &FieldError{FieldRequired}},

		// URL Format
		{VEmail{Required: true}, "google@gmail.com", nil},
		{VEmail{Required: true}, "asdf123@google.com", nil},
		{VEmail{Required: true}, "google@@gmail.com", &FieldError{FieldNoEmail}},
		{VEmail{Required: true}, "some string", &FieldError{FieldNoEmail}},
		{VEmail{Required: true}, "some@string", &FieldError{FieldNoEmail}},
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
