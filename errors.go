package validator

import (
	"strings"
)

type RulesError map[string]error

func (re RulesError) Error() string {
	if len(re) > 0 {
		return "The form contains errors"
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
type FieldErrorSet []error

func (fes *FieldErrorSet) Error() string {
	return "" //????
}

func (fes *FieldErrorSet) HasErrors() bool {
	if len(*fes) == 0 {
		return false
	}

	for _, v := range *fes {
		if v != nil {
			return true
		}
	}

	return false
}

func NewFieldErrorSet(m ...error) *FieldErrorSet {
	res := FieldErrorSet{}
	res = append(res, m...)

	return &res
}
