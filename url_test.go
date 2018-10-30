package validator

import (
	"reflect"
	"testing"
)

func TestVUrl(t *testing.T) {
	tables := []struct {
		validator VUrl
		v         interface{}
		res       *VFieldResult
	}{
		{VUrl{}, nil, nil},

		// Required
		{VUrl{Required: true}, nil, &VFieldResult{FieldRequired}},
		{VUrl{Required: true}, "", &VFieldResult{FieldRequired}},

		// URL Format
		{VUrl{Required: true}, "http://google.com", nil},
		{VUrl{Required: true}, "https://google.com", nil},
		{VUrl{Required: true}, "https://www.google.com/analytics", nil},
		{VUrl{Required: true}, "htps://www.google.com/analytics", &VFieldResult{FieldNoUrl}},
		{VUrl{Required: true}, "some string", &VFieldResult{FieldNoUrl}},
	}

	for _, table := range tables {
		if res := table.validator.CheckValue(table.v); (table.res != res && (table.res == nil || res == nil)) || (table.res != nil && res != nil && !reflect.DeepEqual(*res, *table.res)) {
			t.Errorf("Text validator %+v for \"%v\" got:  \"%v\", want: \"%v\".", table.validator, table.v, res, table.res)
		}
	}
}
