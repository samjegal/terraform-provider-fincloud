package clients

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/common"
	compute "github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/compute/client"
	location "github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/location/client"
	network "github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/network/client"
	sslvpn "github.com/samjegal/terraform-provider-fincloud/fincloud/internal/services/sslvpn/client"
)

type Client struct {
	StopContext context.Context
	Account     *ResourceManagerAccount

	Compute  *compute.Client
	Network  *network.Client
	Location *location.Client
	SslVpn   *sslvpn.Client
}

func (client *Client) Build(ctx context.Context, o *common.ClientOptions) error {
	autorest.Count429AsRetry = false

	client.StopContext = ctx

	client.Compute = compute.NewClient(o)
	client.Network = network.NewClient(o)
	client.Location = location.NewClient(o)
	client.SslVpn = sslvpn.NewClient(o)

	return nil
}
