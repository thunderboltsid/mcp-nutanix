package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Cluster defines the Cluster resource template
func Cluster() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeCluster))+"{uuid}",
		string(ResourceTypeCluster),
		mcp.WithTemplateDescription("Cluster resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// ClusterHandler implements the handler for the Cluster resource
func ClusterHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeCluster, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the Cluster
		return client.V3().GetCluster(ctx, uuid)
	})
}
