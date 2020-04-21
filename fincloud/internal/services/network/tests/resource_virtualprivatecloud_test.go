package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/samjegal/fincloud-sdk-for-go/services/network"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/acceptance"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/clients"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/utils"
)

func TestAccResourceVirtualPrivateCloud_basic(t *testing.T) {
	data := acceptance.BuildTestData(t, "fincloud_virtual_private_cloud", "test")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckResourceVirtualPrivateCloudDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceVirtualPrivateCloud_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceVirtualPrivateCloudExists(data.ResourceName),
					resource.TestCheckResourceAttr(data.ResourceName, "cidr_block", "172.31.0.0/16"),
				),
			},
			{
				ResourceName:      data.ResourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceVirtualPrivateCloud_disappears(t *testing.T) {
	data := acceptance.BuildTestData(t, "fincloud_virtual_private_cloud", "test")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acceptance.PreCheck(t) },
		Providers:    acceptance.SupportedProviders,
		CheckDestroy: testCheckResourceVirtualPrivateCloudDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceVirtualPrivateCloud_basic(data),
				Check: resource.ComposeTestCheckFunc(
					testCheckResourceVirtualPrivateCloudExists(data.ResourceName),
				),
			},
		},
	})
}

func testCheckResourceVirtualPrivateCloudExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		vpcName := rs.Primary.Attributes["name"]
		client := acceptance.FincloudProvider.Meta().(*clients.Client).Network.VirtualPrivateCloudClient
		ctx := acceptance.FincloudProvider.Meta().(*clients.Client).StopContext

		resp, err := client.List(ctx)
		if err != nil {
			return fmt.Errorf("VPC 정보를 가져오는데 실패했습니다. err: %s", err)
		}

		if resp.StatusCode == http.StatusNotFound {
			return fmt.Errorf("VPC 정보가 존재하지 않습니다. err: %s", err)
		}

		for _, content := range *resp.Content {
			if vpcName == *content.VpcName {
				return nil
			}
		}

		return fmt.Errorf("해당 VPC 정보가 존재하지 않습니다. name: %s", vpcName)
	}
}

func testResourceExpandVirtualPrivateCloud(rs *terraform.ResourceState) network.VirtualPrivateCloudParameter {
	vpcName := rs.Primary.Attributes["name"]
	cidrBlock := rs.Primary.Attributes["cidr_block"]

	return network.VirtualPrivateCloudParameter{
		VpcName:  utils.String(vpcName),
		Ipv4Cidr: utils.String(cidrBlock),
	}
}

func testCheckResourceVirtualPrivateCloudDestroy(s *terraform.State) error {
	client := acceptance.FincloudProvider.Meta().(*clients.Client).Network.VirtualPrivateCloudClient
	ctx := acceptance.FincloudProvider.Meta().(*clients.Client).StopContext

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "fincloud_virtual_private_cloud" {
			continue
		}

		name := rs.Primary.Attributes["name"]

		resp, err := client.List(ctx)
		if err != nil {
			return nil
		}

		for _, content := range *resp.Content {
			if name == *content.VpcName {
				return fmt.Errorf("VPC 정보가 삭제되지 않고 존재합니다. name: %s", name)
			}
		}
	}

	return nil
}

func testAccResourceVirtualPrivateCloud_basic(data acceptance.TestData) string {
	return fmt.Sprintf(`
resource "fincloud_virtual_private_cloud" "test" {
	name  = "vpctest-%d"
	cidr_block = "172.31.0.0/16"
}
`, data.RandomInteger)
}
