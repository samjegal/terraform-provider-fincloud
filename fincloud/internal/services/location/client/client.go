package client

import (
	"github.com/samjegal/fincloud-sdk-for-go/services/location"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/common"
)

type Client struct {
	ZoneClient   *location.ZoneClient
	RegionClient *location.RegionClient
}

func NewClient(o *common.ClientOptions) *Client {
	ZoneClient := location.NewZoneClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&ZoneClient.Client, o.ResourceManagerAuthorizer)

	RegionClient := location.NewRegionClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&RegionClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		ZoneClient:   &ZoneClient,
		RegionClient: &RegionClient,
	}
}
