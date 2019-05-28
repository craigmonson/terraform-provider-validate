data "validate" "exact" {
  val   = "foo"
  exact = "foo"
}

data "validate" "one_of" {
  val    = "foo"
  one_of = ["foo", "bar", "baz"]
}

data "validate" "regex" {
  val   = "foo"
  regex = "f.."
}
