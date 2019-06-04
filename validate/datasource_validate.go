package validate

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"regexp"
	"strings"
)

func dataSourceValidateSchema() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceTest,

		Schema: map[string]*schema.Schema{
			"val": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Value to validate",
			},
			"exact": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Value must match exactly",
				ConflictsWith: []string{"one_of", "regex", "not_one_of", "not_regex"},
			},
			"not_exact": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "Value must NOT match exactly",
				ConflictsWith: []string{"exact", "one_of", "not_one_of"},
			},
			"one_of": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of acceptable values",
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Optional: false,
				},
				ConflictsWith: []string{"exact", "not_exact", "not_one_of", "regex", "not_regex"},
			},
			"not_one_of": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of UN-acceptable values",
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Optional: false,
				},
				ConflictsWith: []string{"not_exact", "exact"},
			},
			"regex": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "RE2 regular expression string (see: https://golang.org/s/re2syntax)",
				ConflictsWith: []string{"one_of", "exact"},
			},
			"not_regex": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "RE2 regular expression string to NOT match (see: https://golang.org/s/re2syntax)",
				ConflictsWith: []string{"exact", "not_exact", "one_of"},
			},
		},
	}
}

func dataSourceTest(d *schema.ResourceData, meta interface{}) error {
	check_types, err := getCheckTypes(d)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, c_type := range check_types {
		switch c_type {
		case "exact":
			err := checkExact(d.Get("val").(string), d.Get("exact").(string))
			if err != nil {
				errs = append(errs, err)
			}
		case "not_exact":
			err := checkNotExact(d.Get("val").(string), d.Get("not_exact").(string))
			if err != nil {
				errs = append(errs, err)
			}
		case "one_of":
			err := checkOneOf(d.Get("val").(string), d.Get("one_of").([]interface{}))
			if err != nil {
				errs = append(errs, err)
			}
		case "not_one_of":
			err := checkNotOneOf(d.Get("val").(string), d.Get("not_one_of").([]interface{}))
			if err != nil {
				errs = append(errs, err)
			}
		case "regex":
			err := checkRegex(d.Get("val").(string), d.Get("regex").(string))
			if err != nil {
				errs = append(errs, err)
			}
		case "not_regex":
			err := checkNotRegex(d.Get("val").(string), d.Get("not_regex").(string))
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	if len(errs) != 0 {
		errStrs := []string{}
		for _, e := range errs {
			errStrs = append(errStrs, e.Error())
		}

		return fmt.Errorf("%s", strings.Join(errStrs, ";"))
	}

	return nil
}

var checkExact = func(val, check string) error {
	if val != check {
		return fmt.Errorf("val: '%s' does not match '%s' for exact", val, check)
	}

	return nil
}

var checkNotExact = func(val, check string) error {
	if val == check {
		return fmt.Errorf("val: '%s' matched '%s' for not_exact", val, check)
	}

	return nil
}

var checkOneOf = func(val string, list []interface{}) error {
	for _, c := range list {
		if val == c.(string) {
			return nil
		}
	}

	return fmt.Errorf("val '%s' is not in one_of list: %v", val, list)
}

var checkNotOneOf = func(val string, list []interface{}) error {
	for _, c := range list {
		if val == c.(string) {
			return fmt.Errorf("val '%s' is in not_one_of list: %v", val, list)
		}
	}

	return nil
}

var checkRegex = func(val, pattern string) error {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	if reg.MatchString(val) {
		return nil
	}

	return fmt.Errorf("val '%s' does not match the regex '%s'", val, pattern)
}

var checkNotRegex = func(val, pattern string) error {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	if reg.MatchString(val) {
		return fmt.Errorf("val '%s' matched not_regex '%s'", val, pattern)
	}

	return nil
}

func getCheckTypes(d *schema.ResourceData) ([]string, error) {
	var c_types []string

	if _, ok := d.GetOk("exact"); ok {
		c_types = append(c_types, "exact")
	}
	if _, ok := d.GetOk("not_exact"); ok {
		c_types = append(c_types, "not_exact")
	}
	if _, ok := d.GetOk("one_of"); ok {
		c_types = append(c_types, "one_of")
	}
	if _, ok := d.GetOk("not_one_of"); ok {
		c_types = append(c_types, "not_one_of")
	}
	if _, ok := d.GetOk("regex"); ok {
		c_types = append(c_types, "regex")
	}
	if _, ok := d.GetOk("not_regex"); ok {
		c_types = append(c_types, "not_regex")
	}

	if len(c_types) == 0 {
		return c_types, fmt.Errorf("Must choose at least one attribute: exact, not_exact, one_of, not_one_of, regex, or not_regex (see documentation for acceptable combinations)")
	}

	return c_types, nil
}
