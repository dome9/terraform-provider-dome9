provider "dome9" {
  dome9_access_id  = ""
  dome9_secret_key = ""
}

resource "dome9_iplist" "iplist" {
  name        = "ipList test name"
  description = "this is the descrption of my iplist"

  items {
    ip      = "1.1.4.4/32"
    comment = "test-ip1"
  }
  items {
    ip      = "1.1.1.5/32"
    comment = "test-ip2"
  }
  items {
    ip      = "1.1.5.53/32"
    comment = "test-ip3"
  }
}

data "dome9_iplist" "data" {
  id = 0000
}

output "getId" {
  value = "${data.dome9_iplist.data.id}"
}

output "getDescription" {
  value = "${data.dome9_iplist.data.description}"
}

output "getItems" {
  value = "${data.dome9_iplist.data.items}"
}
