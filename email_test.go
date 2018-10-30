package validator

import (
	"reflect"
	"testing"
)

func TestVEmail(t *testing.T) {
	tables := []struct {
		validator VEmail
		v         interface{}
		res       *VFieldResult
	}{
		{VEmail{}, nil, nil},

		// Required
		{VEmail{Required: true}, nil, &VFieldResult{FieldRequired}},
		{VEmail{Required: true}, "", &VFieldResult{FieldRequired}},

		// URL Format
		{VEmail{Required: true}, "google@gmail.com", nil},
		{VEmail{Required: true}, "asdf123@google.com", nil},
		{VEmail{Required: true}, "google@@gmail.com", &VFieldResult{FieldNoEmail}},
		{VEmail{Required: true}, "some string", &VFieldResult{FieldNoEmail}},
		{VEmail{Required: true}, "some@string", &VFieldResult{FieldNoEmail}},
	}

	for _, table := range tables {
		if res := table.validator.CheckValue(table.v); (table.res != res && (table.res == nil || res == nil)) || (table.res != nil && res != nil && !reflect.DeepEqual(*res, *table.res)) {
			t.Errorf("Text validator %+v for \"%v\" got:  \"%v\", want: \"%v\".", table.validator, table.v, res, table.res)
		}
	}
}
