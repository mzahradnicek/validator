package validator

import (
	"errors"
	"reflect"
)

const (
	FieldWrongType = "ValidatorWrongType" // field has wrong type
	FieldRequired  = "ValidatorRequired"  // field is required

	// Text
	FieldTextMinVal = "ValidatorTextMinVal" // text field length is lower than min
	FieldTextMaxVal = "ValidatorTextMaxVal" // text field length is higher than max

	// Numeric
	FieldNoNumeric          = "ValidatorNumNotNum"       // field is not numeric
	FieldIsNegative         = "ValidatorNumIsNegative"   // numeric field is negative
	FieldIsFloat            = "ValidatorNumIsFloat"      // numeric is float
	FieldHasTooManyDecimals = "ValidatorNumManyDecimals" // field has too many decimals
	FieldNumMinVal          = "ValidatorNumMinVal"       // numeric value is lower than min
	FieldNumMaxVal          = "ValidatorNumMaxVal"       // numeric value is higher than max

	// Date
	FieldNoDate     = "ValidatorDateNotDate"
	MinNoDate       = "ValidatorDateMinNotDate"
	MaxNoDate       = "ValidatorDateMaxNotDate"
	FieldDateMinVal = "ValidatorDateMinVal"
	FieldDateMaxVal = "ValidatorDateMaxVal"

	// Email
	FieldNoEmail = "ValidatorEmailNotEmail"

	// Url
	FieldNoUrl = "ValidatorUrlNotUrl"

	// Regexp
	FieldNoMatch   = "ValidatorRegexpNoMatch"
	FieldBadRegexp = "ValidatorRegexpBad"
)

type VRule interface {
	IsRequired() bool
	CheckValue(v interface{}) error
}

type VRules map[string]VRule

type VResults map[string]error

func (vr VRules) Validate(input interface{}) error {
	ivr := reflect.ValueOf(input)

	switch ivr.Kind() {
	case reflect.Map:
		if mp, ok := input.(map[string]interface{}); ok {
			return vr.ValidateMap(mp)
		} else {
			return errors.New("Wrong map structure")
		}
	// case reflect.Struct:
	default:
		return errors.New("Wrong variable type for validate")
	}
}

func (vr VRules) ValidateMap(input map[string]interface{}) error {
	res := RulesError{}

	for k, rule := range vr {
		val, ok := input[k]

		// check required and not set values
		if !ok || val == nil {
			if rule.IsRequired() {
				res[k] = &FieldErrorSet{&FieldError{FieldRequired}}
			}

			continue
		}

		reflectVal := reflect.ValueOf(val)
		switch reflectVal.Kind() {
		case reflect.Array, reflect.Slice:
			itemsCnt := reflectVal.Len()

			if itemsCnt == 0 && rule.IsRequired() {
				res[k] = &FieldErrorSet{&FieldError{FieldRequired}}
				break
			}

			fes := FieldErrorSet{}

			for i := 0; i < itemsCnt; i++ {
				if err := rule.CheckValue(reflectVal.Index(i).Interface()); err != nil {
					if fe, ok := err.(*FieldError); ok {
						fes = append(fes, fe)
					} else {
						res[k] = err
					}
				}
			}

			if len(fes) > 0 {
				res[k] = &fes
			}
		default:
			if err := rule.CheckValue(val); err != nil {
				if fe, ok := err.(*FieldError); ok {
					res[k] = &FieldErrorSet{fe}
				} else {
					res[k] = err
				}
			}
		}
	}

	if len(res) > 0 {
		return res
	}

	return nil
}

func Validate(input map[string]interface{}, rules VRules) error {
	return rules.Validate(input)
}
