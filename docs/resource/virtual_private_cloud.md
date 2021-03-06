---
subcategory: "Network"
layout: "fincloud"
page_title: "Financial Cloud Resource Manager: fincloud_virtual_private_cloud"
description: |-
  설명
---

# fincloud_virtual_network_cloud

설명 위와 동일

## Example Usage

```hcl
resource "azurerm_resource_group" "example" {
  name     = "acceptanceTestResourceGroup1"
  location = "West US"
}

resource "azurerm_network_security_group" "example" {
  name                = "acceptanceTestSecurityGroup1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
}

resource "azurerm_network_ddos_protection_plan" "example" {
  name                = "ddospplan1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
}

resource "azurerm_virtual_network" "example" {
  name                = "virtualNetwork1"
  location            = azurerm_resource_group.example.location
  resource_group_name = azurerm_resource_group.example.name
  address_space       = ["10.0.0.0/16"]
  dns_servers         = ["10.0.0.4", "10.0.0.5"]

  ddos_protection_plan {
    id     = azurerm_network_ddos_protection_plan.example.id
    enable = true
  }

  subnet {
    name           = "subnet1"
    address_prefix = "10.0.1.0/24"
  }

  subnet {
    name           = "subnet2"
    address_prefix = "10.0.2.0/24"
  }

  subnet {
    name           = "subnet3"
    address_prefix = "10.0.3.0/24"
    security_group = azurerm_network_security_group.example.id
  }

  tags = {
    environment = "Production"
  }
}
```

## Argument Reference

The following arguments are supported:

- `name` - (Required) The name of the virtual network. Changing this forces a
  new resource to be created.

- `ddos_protection_plan` - (Optional) A `ddos_protection_plan` block as documented below.

- `tags` - (Optional) A mapping of tags to assign to the resource.

---

A `ddos_protection_plan` block supports the following:

- `id` - (Required) The Resource ID of DDoS Protection Plan.

- `enable` - (Required) Enable/disable DDoS Protection Plan on Virtual Network.

## Attributes Reference

The following attributes are exported:

- `id` - The virtual NetworkConfiguration ID.

- `name` - The name of the virtual network.

- `resource_group_name` - The name of the resource group in which to create the virtual network.

- `subnet`- One or more `subnet` blocks as defined below.

---

The `subnet` block supports:

- `name` - (Required) The name of the subnet.

- `address_prefix` - (Required) The address prefix to use for the subnet.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

- `create` - (Defaults to 30 minutes) Used when creating the Virtual Network.
- `update` - (Defaults to 30 minutes) Used when updating the Virtual Network.
- `read` - (Defaults to 5 minutes) Used when retrieving the Virtual Network.
- `delete` - (Defaults to 30 minutes) Used when deleting the Virtual Network.

## Import

Virtual Networks can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_virtual_network.exampleNetwork /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Network/virtualNetworks/myvnet1
```
