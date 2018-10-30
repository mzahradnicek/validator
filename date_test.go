package validator

import (
	"reflect"
	"testing"
)

func TestVDate(t *testing.T) {
	tables := []struct {
		validator VDate
		v         interface{}
		res       *VFieldResult
	}{
		{VDate{}, nil, nil},

		// Required
		{VDate{Required: true}, nil, &VFieldResult{FieldRequired}},
		{VDate{Required: true}, "", &VFieldResult{FieldRequired}},

		// write some cases
	}

	for _, table := range tables {
		if res := table.validator.CheckValue(table.v); (table.res != res && (table.res == nil || res == nil)) || (table.res != nil && res != nil && !reflect.DeepEqual(*res, *table.res)) {
			t.Errorf("Text validator %+v for \"%v\" got:  \"%v\", want: \"%v\".", table.validator, table.v, res, table.res)
		}
	}
}
