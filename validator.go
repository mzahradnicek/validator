package validator

import (
	"encoding/json"
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
	CheckValue(v interface{}) *VFieldResult
}

type VRules map[string]VRule

type VFieldResult []string

type VFieldResultSet []*VFieldResult

type VResults map[string]VFieldResultSet

func Validate(input map[string]interface{}, rules VRules) (VResults, bool) {
	res := make(VResults)
	resOk := true

	for k, rule := range rules {
		val, ok := input[k]

		// check required and not set values
		if !ok || val == nil {
			if rule.IsRequired() {
				res[k] = VFieldResultSet{&VFieldResult{FieldRequired}}
				resOk = false
			}

			continue
		}

		reflectVal := reflect.ValueOf(val)
		switch reflectVal.Kind() {
		case reflect.Array, reflect.Slice:
			fres := VFieldResultSet{}
			itemsCnt := reflectVal.Len()

			// check required empty array
			if itemsCnt > 0 {
				setRes := false
				for i := 0; i < itemsCnt; i++ {
					itemErr := rule.CheckValue(reflectVal.Index(i).Interface())
					if itemErr != nil {
						setRes = true
					}

					fres = append(fres, itemErr)
				}

				if setRes {
					res[k] = fres
					resOk = false
				}

			} else if rule.IsRequired() {
				fres = append(fres, &VFieldResult{FieldRequired})
				res[k] = fres
				resOk = false
			}

		default:
			if err := rule.CheckValue(val); err != nil {
				res[k] = VFieldResultSet{err}
				resOk = false
			}
		}

	}

	return res, resOk
}

func (frs VFieldResultSet) MarshalJSON() ([]byte, error) {
	if len(frs) == 1 {
		return json.Marshal(frs[0])
	}

	return json.Marshal(frs)
}
