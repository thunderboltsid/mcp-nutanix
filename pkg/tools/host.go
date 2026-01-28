package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Host defines the Host tool
func HostList() mcp.Tool {
	return mcp.NewTool("host_list",
		mcp.WithDescription("List host resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// HostListHandler implements the handler for the Host list tool
func HostListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeHost,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			hosts, err := client.Converged().Clusters.ListAllHosts(ctx)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"hosts": hosts,
				"count": len(hosts),
			}, nil
		},
	)
}

// HostCount defines the Host count tool
func HostCount() mcp.Tool {
	return mcp.NewTool("host_count",
		mcp.WithDescription("Count host resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// HostCountHandler implements the handler for the Host count tool
func HostCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeHost,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			hosts, err := client.Converged().Clusters.ListAllHosts(ctx)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"resource_type": "Host",
				"count":         len(hosts),
			}, nil
		},
	)
}
