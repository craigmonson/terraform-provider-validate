data "validate_str" "exact" {
  val   = "foo"
  exact = "foo"
}

data "validate_str" "one_of" {
  val    = "foo"
  one_of = ["foo", "bar", "baz"]
}

data "validate_str" "regex" {
  val   = "foo"
  regex = "f.."
}
