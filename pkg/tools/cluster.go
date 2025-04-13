package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Cluster defines the Cluster tool
func ClusterList() mcp.Tool {
	return mcp.NewTool("cluster_list",
		mcp.WithDescription("List cluster resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ClusterListHandler implements the handler for the Cluster list tool
func ClusterListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeCluster,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllCluster(ctx, "")

		},
	)
}

// ClusterCount defines the Cluster count tool
func ClusterCount() mcp.Tool {
	return mcp.NewTool("cluster_count",
		mcp.WithDescription("Count cluster resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ClusterCountHandler implements the handler for the Cluster count tool
func ClusterCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeCluster,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllCluster(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "Cluster",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
