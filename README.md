# Terraform Provider - Validate

## Maintainers

This provider plugin is maintained by Craig Monson

## Requirements

  * [Terraform](https://www.terraform.io/downloads.html) 0.11.x or greater
  * [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)

## Usage

To install, as this is a 3rd party plugin, you'll need to follow the directions
[here](https://www.terraform.io/docs/configuration/providers.html#third-party-plugins).

```hcl
provider "validate" {}

variable "test_exact" {
  type    = "string"
  default = "test_exact"
}

variable "test_one_of" {
  type    = "string"
  default = "test_one_of"
}

variable "test_regex" {
  type    = "string"
  default = "^test_.*"
}

data "validate" "exact" {
  val   = "${var.test_exact}"
  exact = "test_exact"
}

data "validate" "one_of" {
  val    = "${var.test_one_of}"
  one_of = ["test_one_of"]
}

data "validate" "regex" {
  val   = "${var.test_regex}"
  regex = ["test_regex"]
}
```

The `validate` data source will validate input values against a validation check.  IWhen
`terraform plan` is run, the values will be validated against the check, and will raise
an error if it fails.  This allows for ensuring any data requirements (tagging for
example) meet a desired set of criteria.

### Argument Reference

The following arguments are supported:

  * `val` - (Required) The value to be checked.  This should be a variable, and is expected to be a string.
  * `exact` - (Optional) (a string) Check `val` is an exact match of this argument.  `exact`, `one_of` and `regex` are mutually exclusive, but one must exist in the data source.
  * `one_of` - (Optional) (a list) Check `val` is an exact match of one of the items in this list.  `exact`, `one_of` and `regex` are mutually exclusive, but one must exist in the data source.
  * `regex` - (Optional) (a string) Check `val` matches the regular expression expressed as this string.  This utilizes the RE2 regular expression syntax, which can be found [here](https://golang.org/s/re2syntax). `exact`, `one_of` and `regex` are mutually exclusive, but one must exist in the data source.

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
which will execute a terraform template with passinchecks.  Similarly, to execute a
terraform template with _failing_ checks, run:
```
make tf-test-fail
````
