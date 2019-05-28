package test_helpers

import (
	"github.com/hashicorp/terraform/helper/schema"
	"testing"
)

func GetResourceData(t *testing.T, m map[string]interface{}, res *schema.Resource) *schema.ResourceData {
	return schema.TestResourceDataRaw(t, res.Schema, m)
}
