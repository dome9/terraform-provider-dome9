module github.com/terraform-providers/terraform-provider-dome9

go 1.13

//replace github.com/dome9/dome9-sdk-go v1.12.0 => ./../dome9-sdk-go

require (
	github.com/dome9/dome9-sdk-go v1.13.0
	github.com/hashicorp/terraform-plugin-sdk v1.1.0
)
