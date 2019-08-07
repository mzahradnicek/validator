package validator

import (
	"encoding/json"
	"strings"
)

type RulesError map[string]error

func (re RulesError) Error() string {
	if len(re) > 0 {
		return "There are errors"
	} else {
		return ""
	}
}

func (re RulesError) AddFieldError(key string, e error) {
	re[key] = e
}

func NewRulesError() RulesError {
	return RulesError{}
}

/* Field Error handler */
type FieldError []string

func (fe *FieldError) Error() string {
	return strings.Join(*fe, ",")
}

/* List of errors for field */
type FieldErrorSet []*FieldError

func (fes *FieldErrorSet) Error() string {
	return "" //????
}

func (fes *FieldErrorSet) HasError() bool {
	return len(*fes) > 0
}

func (fes *FieldErrorSet) Add(m ...*FieldError) {
	*fes = append(*fes, m...)
}

func (fes FieldErrorSet) MarshalJSON() ([]byte, error) {
	if len(fes) == 1 {
		return json.Marshal(fes[0])
	}

	return json.Marshal(fes)
}

func NewFieldErrorSet(m ...*FieldError) *FieldErrorSet {
	res := FieldErrorSet{}
	res = append(res, m...)

	return &res
}
