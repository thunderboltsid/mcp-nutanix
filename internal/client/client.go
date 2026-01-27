package client

import (
	"github.com/nutanix-cloud-native/prism-go-client/environment"
	"github.com/nutanix-cloud-native/prism-go-client/environment/providers/local"
	"github.com/nutanix-cloud-native/prism-go-client/environment/providers/mcp"
	envtypes "github.com/nutanix-cloud-native/prism-go-client/environment/types"
	prismclientv4 "github.com/nutanix-cloud-native/prism-go-client/v4"
	"k8s.io/klog"
)

var (
	prismClient *NutanixClient
)

func Init(modelcontextclient mcp.ModelContextClient) {
	prismClient = &NutanixClient{
		env:           environment.NewEnvironment(local.NewProvider(), mcp.NewProvider(modelcontextclient)),
		v4ClientCache: prismclientv4.NewClientCache(),
	}
}

func GetPrismClient() *NutanixClient {
	if prismClient == nil {
		panic("Prism client not initialized. Call Init() first.")
	}

	return prismClient
}

type NutanixClient struct {
	env           envtypes.Environment
	v4ClientCache *prismclientv4.ClientCache
}

// V4 returns the v4 converged client
func (n *NutanixClient) V4() *prismclientv4.Client {
	c, err := n.v4ClientCache.GetOrCreate(n)
	if err != nil {
		panic(err)
	}

	return c
}

// Key returns the constant client name
// This implements the CachedClientParams interface of prism-go-client
func (n *NutanixClient) Key() string {
	return "mcp-server"
}

// ManagementEndpoint returns the management endpoint of the Nutanix cluster
// This implements the CachedClientParams interface of prism-go-client
func (n *NutanixClient) ManagementEndpoint() envtypes.ManagementEndpoint {
	mgmtEndpoint, err := n.env.GetManagementEndpoint(envtypes.Topology{})
	if err != nil {
		klog.Errorf("failed to get management endpoint: %s", err.Error())
		return envtypes.ManagementEndpoint{}
	}

	return *mgmtEndpoint
}
