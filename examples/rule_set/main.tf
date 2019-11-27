resource "dome9_rule_set" "ruleset" {
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

data "dome9_rule_set" "data1" {
  id = dome9_rule_set.ruleset.id
}

output "getId" {
  value = "${data.dome9_rule_set.data1.id}"
}

output "getDescription" {
  value = "${data.dome9_rule_set.data1.description}"
}
