package client

import (
	"github.com/samjegal/fincloud-sdk-for-go/services/loadbalancer"
	"github.com/samjegal/terraform-provider-fincloud/fincloud/internal/common"
)

type Client struct {
	LoadBalancerClient         *loadbalancer.Client
	LoadBalancerListenerClient *loadbalancer.ListenerClient
	LoadBalancerServerClient   *loadbalancer.ServerClient
}

func NewClient(o *common.ClientOptions) *Client {
	LoadBalancerClient := loadbalancer.NewClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&LoadBalancerClient.Client, o.ResourceManagerAuthorizer)

	LoadBalancerListenerClient := loadbalancer.NewListenerClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&LoadBalancerListenerClient.Client, o.ResourceManagerAuthorizer)

	LoadBalancerServerClient := loadbalancer.NewServerClientWithBaseURI(o.ResourceManagerEndpoint)
	o.ConfigureClient(&LoadBalancerServerClient.Client, o.ResourceManagerAuthorizer)

	return &Client{
		LoadBalancerClient:         &LoadBalancerClient,
		LoadBalancerListenerClient: &LoadBalancerListenerClient,
		LoadBalancerServerClient:   &LoadBalancerServerClient,
	}
}
