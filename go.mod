module github.com/samjegal/terraform-provider-fincloud

go 1.15

replace github.com/Azure/go-autorest/autorest => github.com/samjegal/go-autorest/autorest v0.11.5-0.20200906111652-5e286818fa7f

require (
	github.com/Azure/go-autorest/autorest v0.11.4
	github.com/hashicorp/terraform-config-inspect v0.0.0-20200806211835-c481b8bfa41e // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.15.0
	github.com/samjegal/fincloud-sdk-for-go v1.9.2
	github.com/samjegal/go-fincloud-helpers v0.2.4
)
