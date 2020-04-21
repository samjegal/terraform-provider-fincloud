package tests

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/acceptance"
)

func TestAccDataSourceVirtualPrivateCloud_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "fincloud_virtual_private_cloud", "test")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acceptance.PreCheck(t) },
		Providers: acceptance.SupportedProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceVirtualPrivateCloud_basic(data),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(data.ResourceName, "cidr_block", "172.31.0.0/16"),
				),
			},
		},
	})
}

func testAccDataSourceVirtualPrivateCloud_basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "fincloud_virtual_private_cloud" "test" {
	name = "vpctest-%d"
	cidr_block = "172.31.0.0/16"
}

data "fincloud_virtual_private_cloud" "test" {
	name = fincloud_virtual_private_cloud.test.name
}
`, data.RandomInteger)
}
