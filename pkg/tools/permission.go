package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Permission defines the Permission tool
func Permission() mcp.Tool {
	return mcp.NewTool("permissions",
		mcp.WithDescription("List permission resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// PermissionHandler implements the handler for the Permission tool
func PermissionHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypePermission,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllPermission(ctx, "")

		},
	)
}
