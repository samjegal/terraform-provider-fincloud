## 1.1.0 ()

FEATURES:

- **New Resource:** `fincloud_vpn_connection`: Cannot be created or deleted
- **New Resource:** `fincloud_route_table`

## 1.0.1 (April 22, 2020)

BUG FIXES:

- Fix incorrect setting of home directory path in Windows environment
- Fixed an issue where the parameter value was empty in the initial configuration file
- Changed the zip method of compressing binary for windows in task of github action

## 1.0.0 (April 21, 2020)

NOTE:

- Major Release Version: Version 1.0 of Financial Cloud Provider is a major version.
- Terraform 0.12: You must upgrade to Terraform 0.12 to use version 1.0 of the Financial Cloud Provider

FEATURES:

- **New Resource:** `fincloud_virtual_private_cloud`
- **New Resource:** `fincloud_network_acl`
- **New Resource:** `fincloud_network_acl_rule`
- **New Resource:** `fincloud_subnet`
- **New Resource:** `fincloud_region`
- **New Resource:** `fincloud_zone`
- **New Resource:** `fincloud_server`
- **New Resource:** `fincloud_storage`
- **New Resource:** `fincloud_network_interface`
- **New Resource:** `fincloud_public_ip`
- **New Resource:** `fincloud_security_group`
- **New Resource:** `fincloud_security_group_rule`
- **New Resource:** `fincloud_init_script`
- **New Resource:** `fincloud_login_key`

IMPROVEMENTS:

- Change the overall directory structure with reference to [terraform-provider-azurerm](https://github.com/terraform-providers/terraform-provider-azurerm)
