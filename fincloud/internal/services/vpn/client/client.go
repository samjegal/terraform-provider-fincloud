package client

import (
	"github.com/samjegal/fincloud-sdk-for-go/services/sslvpn"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/common"
)

type Client struct {
	VpnConnectionClient *sslvpn.Client
}

func NewClient(o *common.ClientOptions) *Client {
	VpnConnectionClient := sslvpn.NewClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&VpnConnectionClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		VpnConnectionClient: &VpnConnectionClient,
	}
}
