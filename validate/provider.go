package validate

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"validate_str": dataSourceValidateStr(),
			//"validate_int": dataSourceValidateInt(),
		},
	}
}
