provider "dome9" {
  dome9_access_id  = ""
  dome9_secret_key = ""
}

resource "dome9_iplist" "iplist" {
  name        = "sendbox"
  description = "DESC"

  items = [
    {
      ip      = "1.1.4.4"
      comment = "test-ip1"
    },
    {
      ip      = "1.1.1.5"
      comment = "test-ip2"
    },
    {
      ip      = "1.1.5.53"
      comment = "test-ip3"
    },
  ]
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
