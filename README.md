# NOT MAINTAINED
As of Terraform 0.13, variable validation is an included feature, so this project is moot.

# Terraform Provider - Validate
[![Build Status](https://travis-ci.com/craigmonson/terraform-provider-validate.svg?branch=master)](https://travis-ci.com/craigmonson/terraform-provider-validate) [![Coverage Status](https://coveralls.io/repos/github/craigmonson/terraform-provider-validate/badge.svg?branch=master)](https://coveralls.io/github/craigmonson/terraform-provider-validate?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/craigmonson/terraform-provider-validate)](https://goreportcard.com/report/github.com/craigmonson/terraform-provider-validate) [![Release Badge](https://img.shields.io/github/release/craigmonson/terraform-provider-validate.svg)](https://github.com/craigmonson/terraform-provider-validate/releases/latest)

## Maintainers

This provider plugin is maintained by Craig Monson

## Requirements

  * [Terraform](https://www.terraform.io/downloads.html) 0.11.x or greater
  * [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

## Installation

To install, as this is a 3rd party plugin, you'll need to follow the directions
[here](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).

Basically, download the [latest](https://github.com/craigmonson/terraform-provider-validate/releases/latest), decompress it, and place the executable in your
`$HOME/.terraform.d/plugins/<OS>_<ARCH>/` directory (for unix based systems), or `%APPDATA%\terraform.d\plugins\<OS>_<ARCH>` for windows.

The file should be named: `terraform-provider-validate_vX.X.X` for \*nix versions or `terraform-provider-validate_vX.X.X.exe` for windows.  (where X.X.X matches the version) 

`terraform init` will pull from the above directory.

Example: 
download the binary for os-x, unzip, and make sure the file is saved as:
`/Users/<USERNAME>/.terraform.d/plugins/darwin_amd64/terraform-provider-validate_vX.X.X`

Or, on windows (possibly):
`C:\Users\<USERNAME>\AppData\Roaming\terraform.d\plugins\windows_amd64\terraform-provider-validate_vX.X.X.exe`

## Usage

```hcl
provider "validate" {}

variable "test_string" {
  type    = string
  default = "Test String"
}


##             ##
## Validations ##
##             ##

# Pass if the strings are an exact match.
data "validate" "exact" {
  val   = var.test_string
  exact = "Test String"
}

# Pass if the strings do NOT match
data "validate" "not_exact" {
  val       = var.test_string
  not_exact = "This will only pass if var.test_exact doesn't match this"
}

# Pass if var.test_string matches one in the list
data "validate" "one_of" {
  val    = var.test_string
  one_of = ["List", "of", "Possible", "Test String", "s"]
}

# Pass if var.test_string is NOT in the list
data "validate" "not_one_of" {
  val        = var.test_string
  not_one_of = ["Not", "in", "this", "list"]
}

# Pass if the regex is matched (NOTE: Must follow HCL syntax)
data "validate" "regex" {
  val   = var.test_string
  regex = "^Test"
}

# Pass if it does NOT match the regex
data "validate" "not_regex" {
  val       = var.test_string
  not_regex = "^No [m|M]atch"
}

# This validation is optional, so it's ok if val is empty
data "validate" "optional_exact" {
  val      = ""
  exact    = "This variable is optional, so can be empty"
  optional = true
}

# Customize the error message to better describe a failed validation
data "validate" "customized_error_message" {
  val       = var.test_string
  exact     = "Fail me"
  error_msg = "'${var.test_string}' did not match 'Fail me'"
}
```

The `validate` data source will validate input values against a validation check.  When
`terraform plan` is run, the values will be validated against the check, and will raise
an error if it fails.  This allows for ensuring any data requirements (tagging for
example) meet a desired set of criteria.

### Argument Reference

The following arguments are supported:

  * `val` - (Required) (string) The value to be checked.  This should be a variable.
  * `optional` - (Optional) (bool) (Default: false) If set to true, this will allow the check to be optional.  If it's empty, ie: "", then the check will pass, regardless if it passes the underlying check.
  * `error_msg` - (Optional) (string) This message will be displayed instead of the default one if this validation check fails.  This allows more customized error outputs.

At least one of these arguments must also exist, but combinations are possible (see compatability matrix):

  * `exact` - (Optional) (a string) Check `val` is an exact match of this argument.
  * `not_exact` - (Optional) (a string) Check `val` can not match exactly this argument.
  * `one_of` - (Optional) (a list) Check `val` is an exact match of one of the items in this list.
  * `not_one_of` - (Optional) (a list) Check `val` can not exactly match one of the items in this list.
  * `regex` - (Optional) (a string) Check `val` matches the regular expression expressed as this string.  This utilizes the RE2 regular expression syntax, which can be found [here](https://golang.org/s/re2syntax).
  * `not_regex` - (Optional) (a string) Check `val` does not match the regular expression expressed as this string.  This utilizes the RE2 regular expression syntax, which can be found [here](https://golang.org/s/re2syntax).

### Compatibility Matrix

|                  | exact | not\_exact | one\_of | not\_one\_of | regex | not\_regex |
|-----------------:|:-----:|:----------:|:-------:|:------------:|:-----:|:----------:|
| **exact**        |       |      N     |    N    |      N       |   N   |     N      |
| **not\_exact**   |   N   |            |    N    |      N       | **Y** |   **Y**    |
| **one\_of**      |   N   |      N     |         |      N       |   N   |     N      |
| **not\_one\_of** |   N   |      N     |    N    |              | **Y** |   **Y**    |
| **regex**        |   N   |    **Y**   |    N    |    **Y**     |       |   **Y**    |
| **not\_regex**   |   N   |    **Y**   |    N    |    **Y**     | **Y** |            |


### Attributes Reference

There are no attributes for this data source.

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org/)
installed, and a correctly setup [GOPATH](https://golang.org/doc/code.html#GOPATH).

To compile the provider, run:
```
make build
```
This will build the provider, and put the binary in the local directory.

To test, simply run:
```
make test
```

To run the integration tests (check that it works through terraform itself), run:
```
make tf-test
```
which will execute a terraform template with passing checks.  Similarly, to execute a
terraform template with _failing_ checks, run:
```
make tf-test-fail
````
