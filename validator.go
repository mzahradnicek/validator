package validator

import (
	"encoding/json"
)

const (
	FieldWrongType = "ValidatorWrongType" // field has wrong type
	FieldRequired  = "ValidatorRequired"  // field is required
	FieldMinVal    = "ValidatorMinVal"    // field is lower than min
	FieldMaxVal    = "ValidatorMaxVal"    // field is higher than max

	// Numeric
	FieldNoNumeric          = "ValidatorNumNotNum"       // field is not numeric
	FieldIsNegative         = "ValidatorNumIsNegative"   // numeric field is negative
	FieldIsFloat            = "ValidatorNumIsFloat"      // numeric is float
	FieldHasTooManyDecimals = "ValidatorNumManyDecimals" // field has too many decimals

	// Date
	FieldNoDate = "ValidatorDateNotDate"
	MinNoDate   = "ValidatorDateMinNotDate"
	MaxNoDate   = "ValidatorDateMaxNotDate"

	// Email
	FieldNoEmail = "ValidatorEmailNotEmail"

	// Url
	FieldNoUrl = "ValidatorUrlNotUrl"

	// Regexp
	FieldNoMatch   = "ValidatorRegexpNoMatch"
	FieldBadRegexp = "ValidatorRegexpBad"
)

type VRule interface {
	CheckValue(v string) *VFieldResult
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

		if !ok {
			val = ""
		}

		switch val.(type) {
		case string:
			if err := rule.CheckValue(val.(string)); err != nil {
				res[k] = VFieldResultSet{err}
				resOk = false
			}
		case []string:
			fres := VFieldResultSet{}
			setRes := false

			if len(val.([]string)) == 0 {
				val = []string{""}
			}

			for _, x := range val.([]string) {
				err := rule.CheckValue(x)
				if err != nil {
					setRes = true
				}

				fres = append(fres, err)
			}

			if setRes {
				res[k] = fres
				resOk = false
			}

		default:
			res[k] = VFieldResultSet{&VFieldResult{FieldWrongType}}
			resOk = false
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
