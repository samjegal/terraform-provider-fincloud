package client

import (
	"github.com/samjegal/fincloud-sdk-for-go/services/compute"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/common"
)

type Client struct {
	SecurityGroupClient     *compute.SecurityGroupClient
	SecurityGroupRuleClient *compute.SecurityGroupRuleClient
	InboundClient           *compute.InboundRuleClient
	OutboundClient          *compute.OutboundRuleClient

	NetworkInterfaceClient *compute.NetworkInterfaceClient
	PublicIpClient         *compute.PublicIPAddressClient

	ServerClient       *compute.ServerClient
	RootPasswordClient *compute.RootPasswordClient
	StorageClient      *compute.StorageClient

	LoginKeyClient   *compute.LoginKeyClient
	InitScriptClient *compute.InitScriptClient
}

func NewClient(o *common.ClientOptions) *Client {
	SecurityGroupClient := compute.NewSecurityGroupClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&SecurityGroupClient.Client, o.ResourceManagerAuthorizer)

	SecurityGroupRuleClient := compute.NewSecurityGroupRuleClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&SecurityGroupRuleClient.Client, o.ResourceManagerAuthorizer)

	InboundClient := compute.NewInboundRuleClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&InboundClient.Client, o.ResourceManagerAuthorizer)

	OutboundClient := compute.NewOutboundRuleClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&OutboundClient.Client, o.ResourceManagerAuthorizer)

	NetworkInterfaceClient := compute.NewNetworkInterfaceClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&NetworkInterfaceClient.Client, o.ResourceManagerAuthorizer)

	ServerClient := compute.NewServerClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&ServerClient.Client, o.ResourceManagerAuthorizer)

	RootPasswordClient := compute.NewRootPasswordClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&RootPasswordClient.Client, o.ResourceManagerAuthorizer)

	LoginKeyClient := compute.NewLoginKeyClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&LoginKeyClient.Client, o.ResourceManagerAuthorizer)

	InitScriptClient := compute.NewInitScriptClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&InitScriptClient.Client, o.ResourceManagerAuthorizer)

	StorageClient := compute.NewStorageClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&StorageClient.Client, o.ResourceManagerAuthorizer)

	PublicIpClient := compute.NewPublicIPAddressClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&PublicIpClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		SecurityGroupClient:     &SecurityGroupClient,
		SecurityGroupRuleClient: &SecurityGroupRuleClient,
		InboundClient:           &InboundClient,
		OutboundClient:          &OutboundClient,

		NetworkInterfaceClient: &NetworkInterfaceClient,
		PublicIpClient:         &PublicIpClient,

		ServerClient:       &ServerClient,
		RootPasswordClient: &RootPasswordClient,
		StorageClient:      &StorageClient,

		LoginKeyClient:   &LoginKeyClient,
		InitScriptClient: &InitScriptClient,
	}
}
