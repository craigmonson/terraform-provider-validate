package test_helpers

import (
	"github.com/hashicorp/terraform/helper/schema"
	"testing"
)

func GetResourceData(t *testing.T, m map[string]interface{}, res *schema.Resource) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, res.Schema, m)
}

func Contains(list []string, val string) bool {
	for _, c := range list {
		if val == c {
			return true
		}
	}
	return false
}

func NotContains(list []string, val string) bool {
	return !Contains(list, val)
}

var RegexCounter int
var ExactCounter int
var OneOfCounter int

var RegexMocker = func(val, pattern string) error {
	RegexCounter++

	return nil
}

var ExactMocker = func(val, check string) error {
	ExactCounter++

	return nil
}

var OneOfMocker = func(val string, list []interface{}) error {
	OneOfCounter++

	return nil
}
