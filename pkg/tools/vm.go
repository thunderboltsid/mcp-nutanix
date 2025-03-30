package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// VM defines the VM tool
func VM() mcp.Tool {
	return mcp.NewTool("vms",
		mcp.WithDescription("List vm resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// VMHandler implements the handler for the VM tool
func VMHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeVM,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllVM(ctx, "")

		},
	)
}
