package validate

import (
	"fmt"
	th "github.com/craigmonson/terraform-provider-validate/validate/test_helpers"
	"regexp"
	//"github.com/hashicorp/terraform/helper/schema"
	"testing"
)

func TestGetCheckTypesExact(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["exact"] = "bar"
	res_data := th.GetResourceData(t, m, res)
	checks, err := getCheckTypes(res_data)

	if err != nil {
		t.Errorf("Error occurred: %e", err)
	}
	if th.NotContains(checks, "exact") {
		t.Errorf("'exact' not returned")
	}
}

func TestGetCheckTypesNotExact(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["not_exact"] = "bar"
	res_data := th.GetResourceData(t, m, res)
	checks, err := getCheckTypes(res_data)

	if err != nil {
		t.Errorf("Error occurred: %e", err)
	}
	if th.NotContains(checks, "not_exact") {
		t.Errorf("'not_exact' not returned")
	}
}

func TestGetCheckTypesOneOf(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["one_of"] = []string{"bar"}
	res_data := th.GetResourceData(t, m, res)
	checks, err := getCheckTypes(res_data)

	if err != nil {
		t.Errorf("Error occurred: %e", err)
	}
	if th.NotContains(checks, "one_of") {
		t.Errorf("'one_of' not returned")
	}
}

func TestGetCheckTypesNotOneOf(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["not_one_of"] = []string{"bar"}
	res_data := th.GetResourceData(t, m, res)
	checks, err := getCheckTypes(res_data)

	if err != nil {
		t.Errorf("Error occurred: %e", err)
	}
	if th.NotContains(checks, "not_one_of") {
		t.Errorf("'not_one_of' not returned")
	}
}

func TestGetCheckTypesRegex(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["regex"] = "bar"
	res_data := th.GetResourceData(t, m, res)
	checks, err := getCheckTypes(res_data)

	if err != nil {
		t.Errorf("Error occurred: %e", err)
	}
	if th.NotContains(checks, "regex") {
		t.Errorf("'regex' not returned")
	}
}

func TestGetCheckTypesNotRegex(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["not_regex"] = "bar"
	res_data := th.GetResourceData(t, m, res)
	checks, err := getCheckTypes(res_data)

	if err != nil {
		t.Errorf("Error occurred: %e", err)
	}
	if th.NotContains(checks, "not_regex") {
		t.Errorf("'not_regex' not returned")
	}
}

func TestGetCheckTypesNone(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 1)

	m["val"] = "foo"
	res_data := th.GetResourceData(t, m, res)
	checks, err := getCheckTypes(res_data)

	if err == nil {
		t.Errorf("No Error: No exact, one_of, or regex selected, but didn't send error")
	}
	if len(checks) != 0 {
		t.Errorf("checks is not empty? %v", checks)
	}
}

func TestExactPass(t *testing.T) {
	err := checkExact("foo", "foo")

	if err != nil {
		t.Errorf("Exact check failed: 'foo' vs 'foo'")
	}
}

func TestExactFail(t *testing.T) {
	err := checkExact("foo", "bar")

	if err == nil {
		t.Errorf("Exact check did not fail: 'foo' vs 'bar'")
	}
}

func TestNotExactPass(t *testing.T) {
	err := checkNotExact("foo", "bar")

	if err != nil {
		t.Errorf("NotExact check failed: 'foo' vs 'bar'")
	}
}

func TestNotExactFail(t *testing.T) {
	err := checkNotExact("foo", "foo")

	if err == nil {
		t.Errorf("NotExact check failed: 'foo' vs 'foo'")
	}
}

func TestCheckOneOfPass(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["one_of"] = []string{"foo"}
	res_data := th.GetResourceData(t, m, res)

	err := checkOneOf("foo", res_data.Get("one_of").([]interface{}))

	if err != nil {
		t.Error("one_of check failed: 'foo' vs '[foo]'")
	}
}

func TestCheckOneOfFail(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["one_of"] = []string{"bar"}
	res_data := th.GetResourceData(t, m, res)

	err := checkOneOf("foo", res_data.Get("one_of").([]interface{}))

	if err == nil {
		t.Error("one_of check did not fail: 'foo' vs '[bar]'")
	}
}

func TestCheckNotOneOfPass(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["not_one_of"] = []string{"bar"}
	res_data := th.GetResourceData(t, m, res)

	err := checkNotOneOf("foo", res_data.Get("not_one_of").([]interface{}))

	if err != nil {
		t.Error("not_one_of check failed: 'foo' vs '[bar]' should not error")
	}
}

func TestCheckNotOneOfFail(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["not_one_of"] = []string{"foo"}
	res_data := th.GetResourceData(t, m, res)

	err := checkNotOneOf("foo", res_data.Get("not_one_of").([]interface{}))

	if err == nil {
		t.Error("not_one_of check should error: 'foo' vs '[foo]'")
	}
}

func TestCheckRegexPass(t *testing.T) {
	err := checkRegex("foo", "f..")

	if err != nil {
		t.Error("regex check failed: 'foo' vs '/foo/'")
	}
}

func TestCheckRegexFail(t *testing.T) {
	err := checkRegex("foo", "b..")

	if err == nil {
		t.Error("regex check did not fail: 'foo' vs 'b..'")
	}
}

func TestCheckNotRegexPass(t *testing.T) {
	err := checkNotRegex("foo", "b..")

	if err != nil {
		t.Error("not_regex check failed: 'foo' vs 'b..'")
	}
}

func TestCheckNotRegexFail(t *testing.T) {
	err := checkNotRegex("foo", "f..")

	if err == nil {
		t.Error("not_regex check did not fail: 'foo' vs 'f..'")
	}
}

func TestCheckRegexBadRegex(t *testing.T) {
	err := checkRegex("foo", "/[0-9]++/")

	if err == nil {
		t.Error("regex with bad pattern did not fail")
	}
}

func TestCheckNotRegexBadRegex(t *testing.T) {
	err := checkNotRegex("foo", "/[0-9]++/")

	if err == nil {
		t.Error("not_regex with bad pattern did not fail")
	}
}

func TestDataSouceTestExact(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["exact"] = "foo"
	res_data := th.GetResourceData(t, m, res)

	old := checkExact
	checkExact = th.ExactMocker

	_ = dataSourceTest(res_data, nil)
	if th.ExactCounter != 1 {
		t.Error("exact check was not called.")
	}

	checkExact = old
	th.ExactCounter = 0
}

func TestDataSouceTestNotExact(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["not_exact"] = "foo"
	res_data := th.GetResourceData(t, m, res)

	old := checkNotExact
	checkNotExact = th.ExactMocker

	_ = dataSourceTest(res_data, nil)
	if th.ExactCounter != 1 {
		t.Error("not_exact check was not called.")
	}

	checkNotExact = old
	th.ExactCounter = 0
}

func TestDataSouceTestOneOf(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["one_of"] = []string{"foo"}
	res_data := th.GetResourceData(t, m, res)

	old := checkOneOf
	checkOneOf = th.OneOfMocker

	_ = dataSourceTest(res_data, nil)
	if th.OneOfCounter != 1 {
		t.Error("one_of check was not called.")
	}

	checkOneOf = old
	th.OneOfCounter = 0
}

func TestDataSouceTestNotOneOf(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["not_one_of"] = []string{"foo"}
	res_data := th.GetResourceData(t, m, res)

	old := checkNotOneOf
	checkNotOneOf = th.OneOfMocker

	_ = dataSourceTest(res_data, nil)
	if th.OneOfCounter != 1 {
		t.Error("one_of check was not called.")
	}

	checkNotOneOf = old
	th.OneOfCounter = 0
}

func TestDataSouceTestRegex(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["regex"] = "f.."
	res_data := th.GetResourceData(t, m, res)

	old := checkRegex
	checkRegex = th.RegexMocker

	_ = dataSourceTest(res_data, nil)
	if th.RegexCounter != 1 {
		t.Error("regex check was not called.")
	}

	checkRegex = old
	th.RegexCounter = 0
}

func TestDataSouceTestNotRegex(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["not_regex"] = "f.."
	res_data := th.GetResourceData(t, m, res)

	old := checkNotRegex
	checkNotRegex = th.RegexMocker

	_ = dataSourceTest(res_data, nil)
	if th.RegexCounter != 1 {
		t.Error("regex check was not called.")
	}

	checkNotRegex = old
	th.RegexCounter = 0
}

func TestDataSourceTestMultipleErrors(t *testing.T) {
	res := dataSourceValidateSchema()
	m := make(map[string]interface{}, 2)

	m["val"] = "foo"
	m["not_regex"] = "foo"
	m["regex"] = "bar"
	res_data := th.GetResourceData(t, m, res)

	checkErr := dataSourceTest(res_data, nil)
	if checkErr == nil {
		t.Error("errors were not returned")
	}

	reg, err := regexp.Compile("does not match the regex.+matched not_regex")
	if err != nil {
		t.Error(fmt.Errorf("Regex failed: %e", err))
	} else {
		if !reg.MatchString(checkErr.Error()) {
			t.Error(fmt.Errorf("Multiple errors not returned: %s", checkErr.Error()))
		}
	}
}

func TestIsOptionalCheck(t *testing.T) {
	if !isOptionalCheck("", true) {
		t.Error("isOptionalCheck did not return true")
	}
	if isOptionalCheck("", false) {
		t.Error("isOptionalCheck returned true")
	}
}
