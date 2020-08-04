package validator

import (
	"reflect"
	"testing"
)

func TestVUrl(t *testing.T) {
	tables := []struct {
		validator VUrl
		v         interface{}
		res       *FieldError
	}{
		{VUrl{}, nil, nil},

		// Required
		{VUrl{Required: true}, nil, &FieldError{FieldRequired}},
		{VUrl{Required: true}, "", &FieldError{FieldRequired}},

		// URL Format
		{VUrl{Required: true}, "http://google.com", nil},
		{VUrl{Required: true}, "https://google.com", nil},
		{VUrl{Required: true}, "https://www.google.com/analytics", nil},
		{VUrl{Required: true}, "htps://www.google.com/analytics", &FieldError{FieldNoUrl}},
		{VUrl{Required: true}, "some string", &FieldError{FieldNoUrl}},
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
