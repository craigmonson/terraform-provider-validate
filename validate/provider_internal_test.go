package validate

import (
	"github.com/hashicorp/terraform/helper/schema"
	"testing"
)

func TestProviderCompiles(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Errorf("Provider failed to 'compile': %e", err)
	}
}
