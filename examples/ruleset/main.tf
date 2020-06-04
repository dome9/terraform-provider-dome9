resource "dome9_ruleset" "ruleset" {
  name               = "some_ruleset"
  description        = "this is the descrption of my ruleset"
  cloud_vendor       = "aws"
  language           = "en"
  hide_in_compliance = false
  is_template        = false
  rules {
    name           = "some_rule2"
    logic          = "EC2 should x"
    severity       = "High"
    description    = "rule description here"
    compliance_tag = "ct"
    domain         = "test"
    priority       = "high"
    control_title  = "ct"
    rule_id        = ""
    is_default     = false
  }
}

data "dome9_ruleset" "data1" {
  id = dome9_ruleset.ruleset.id
}

output "getId" {
  value = "${data.dome9_ruleset.data1.id}"
}

output "getDescription" {
  value = "${data.dome9_ruleset.data1.description}"
}
