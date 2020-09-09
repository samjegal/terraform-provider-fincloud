package client

import (
	"github.com/samjegal/fincloud-sdk-for-go/services/vpc"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/common"
)

type Client struct {
	Client *vpc.Client
}

func NewClent(o *common.ClientOptions) *Client {
	VpcClient := vpc.NewClient()
	o.ConfigureClient(&VpcClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		Client: &VpcClient,
	}
}
