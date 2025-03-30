package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Host defines the Host tool
func Host() mcp.Tool {
	return mcp.NewTool("hosts",
		mcp.WithDescription("List host resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// HostHandler implements the handler for the Host tool
func HostHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeHost,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Special case for Host which doesn't take a filter
			return client.V3().ListAllHost(ctx)

		},
	)
}
