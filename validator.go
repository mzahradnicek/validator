package validator

import (
	"encoding/json"
)

const (
	FieldWrongType = "WrongType" // field has wrong type
	FieldRequired  = "Required"  // field is required
	FieldMinVal    = "MinVal"    // field is lower than min
	FieldMaxVal    = "MaxVal"    // field is higher than max

	// Numeric
	FieldNoNumeric          = "NumNotNum"       // field is not numeric
	FieldIsNegative         = "NumIsNegative"   // numeric field is negative
	FieldIsFloat            = "NumIsFloat"      // numeric is float
	FieldHasTooManyDecimals = "NumManyDecimals" // field has too many decimals

	// Date
	FieldNoDate = "DateNotDate"
	MinNoDate   = "DateMinNotDate"
	MaxNoDate   = "DateMaxNotDate"

	// Email
	FieldNoEmail = "EmailNotEmail"

	// Url
	FieldNoUrl = "UrlNotUrl"

	// Regexp
	FieldNoMatch   = "RegexpNoMatch"
	FieldBadRegexp = "RegexpBad"
)

type VRule interface {
	CheckValue(v string) *VFieldResult
}

type VRules map[string]VRule

type VFieldResult []string

type VFieldResultSet []*VFieldResult

type VResults map[string]VFieldResultSet

func Validate(input map[string]interface{}, rules VRules) VResults {
	res := make(VResults)
	for k, rule := range rules {

		val, ok := input[k]

		if !ok {
			val = ""
		}

		switch val.(type) {
		case string:
			if err := rule.CheckValue(val.(string)); err != nil {
				res[k] = VFieldResultSet{err}
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
			}

		default:
			res[k] = VFieldResultSet{&VFieldResult{FieldWrongType}}
		}
	}

	return res
}

func (frs VFieldResultSet) MarshalJSON() ([]byte, error) {
	if len(frs) == 1 {
		return json.Marshal(frs[0])
	}

	return json.Marshal(frs)
}
