package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Cluster defines the Cluster tool
func Cluster() mcp.Tool {
	return mcp.NewTool("clusters",
		mcp.WithDescription("List cluster resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ClusterHandler implements the handler for the Cluster tool
func ClusterHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeCluster,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllCluster(ctx, "")

		},
	)
}
