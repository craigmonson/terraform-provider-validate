package validate

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"regexp"
)

func dataSourceValidate() *schema.Resource {
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
				ConflictsWith: []string{"one_of", "regex"},
			},
			"one_of": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of acceptable values",
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Optional: false,
				},
				ConflictsWith: []string{"exact", "regex"},
			},
			"regex": {
				Type:          schema.TypeString,
				Optional:      true,
				Description:   "RE2 regular expression string (see: https://golang.org/s/re2syntax)",
				ConflictsWith: []string{"one_of", "exact"},
			},
		},
	}
}

func dataSourceTest(d *schema.ResourceData, meta interface{}) error {
	check_type, err := getCheckType(d)
	if err != nil {
		return err
	}

	switch check_type {
	case "exact":
		return checkExact(d.Get("val").(string), d.Get("exact").(string))
	case "one_of":
		return checkOneOf(d.Get("val").(string), d.Get("one_of").([]interface{}))
	case "regex":
		return checkRegex(d.Get("val").(string), d.Get("regex").(string))
	}

	return nil
}

func checkExact(val, check string) error {
	if val != check {
		return fmt.Errorf("val: '%s' does not match '%s' for exact", val, check)
	}

	return nil
}

func checkOneOf(val string, list []interface{}) error {
	for _, c := range list {
		if val == c.(string) {
			return nil
		}
	}

	return fmt.Errorf("val '%s' is not in  list: %v", val, list)
}

func checkRegex(val, pattern string) error {
	reg, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	if reg.MatchString(val) {
		return nil
	}

	return fmt.Errorf("val '%s' does not match the regex '%s'", val, pattern)
}

func getCheckType(d *schema.ResourceData) (string, error) {
	if _, ok := d.GetOk("exact"); ok {
		return "exact", nil
	}
	if _, ok := d.GetOk("one_of"); ok {
		return "one_of", nil
	}
	if _, ok := d.GetOk("regex"); ok {
		return "regex", nil
	}

	return "", fmt.Errorf("Must choose one attribute: exact, one_of, or regex")
}
