package client

import (
	"github.com/samjegal/fincloud-sdk-for-go/services/sslvpn"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/common"
)

type Client struct {
	SslVpnConnectionClient *sslvpn.Client
}

func NewClient(o *common.ClientOptions) *Client {
	SslVpnConnectionClient := sslvpn.NewClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&SslVpnConnectionClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		SslVpnConnectionClient: &SslVpnConnectionClient,
	}
}
