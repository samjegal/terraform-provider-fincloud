package client

import (
	"github.com/samjegal/fincloud-sdk-for-go/services/vpc"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/common"
)

type Client struct {
	VpcClient        *vpc.Client
	NetworkACLClient *vpc.NetworkACLClient
}

func NewClent(o *common.ClientOptions) *Client {
	VpcClient := vpc.NewClient()
	o.ConfigureClient(&VpcClient.Client, o.ResourceManagerAuthorizer)

	NetworkACLClient := vpc.NewNetworkACLClient()
	o.ConfigureClient(&NetworkACLClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		VpcClient:        &VpcClient,
		NetworkACLClient: &NetworkACLClient,
	}
}
