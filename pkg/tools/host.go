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
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Special case for Host which doesn't take a filter
			return client.V3().ListAllHost(ctx)

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
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Special case for Host which doesn't take a filter
			resp, err := client.V3().ListAllHost(ctx)

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "Host",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
