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
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			clusters, err := client.Converged().Clusters.List(ctx)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"clusters": clusters,
				"count":    len(clusters),
			}, nil
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
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			clusters, err := client.Converged().Clusters.List(ctx)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"resource_type": "Cluster",
				"count":         len(clusters),
			}, nil
		},
	)
}
