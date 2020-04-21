# Terraform Provider for Financial Cloud

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x +
- [Go](https://golang.org/doc/install) 1.13.x (to build the provider plugin)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

## Installation

This provider is available on macOS and Windows platforms.

### MacOS / Homebrew

```bash
$ brew tap samjegal/fincloud
$ brew install terraform-provider-fincloud
```

### Windows

```powershell
$ choco install terraform-provider-fincloud
```

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is _required_). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

First Clone repository to: `$GOPATH/src/github.com/samjegal/terraform-provider-fincloud`

```bash
$ mkdir -p $GOPATH/src/github.com/samjegal; cd $GOPATH/src/github.com/samjegal
$ git clone git@github.com:samjegal/terraform-provider-fincloud
$ cd $GOPATH/src/github.com/samjegal/terraform-provider-fincloud
```

At this point you can compile the provider by running `make` and put the provider binary in the `$GOPATH/bin` directory.

```bash
$ make
...
$ $GOPATH/bin/terraform-provider-fincloud
```

## Usage Example

```terraform
resource "fincloud_virtual_private_cloud" "example" {
  name       = "dev"
  cidr_block = "172.31.0.0/16"
}

resource "fincloud_network_acl" "example" {
  vpc_id = fincloud_virtual_private_cloud.example.id

  name        = "dev"
  description = "Example Description"
}

resource "fincloud_network_acl_rule" "example" {
  network_acl_id = fincloud_network_acl.example.id

  rule {
    direction   = "inbound"
    protocol    = "tcp"
    cidr_block  = "0.0.0.0/0"
    port        = "4000"
    allow       = true
    priority    = 0
    description = "Inbound Port 4000"
  }
}
```

## Contributing

This repository is inspired by Microsoft Azure Terraform. The code is based on [terraform-provider-azurerm](https://github.com/terraform-providers/terraform-provider-azurerm)
