package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Role defines the Role tool
func Role() mcp.Tool {
	return mcp.NewTool("roles",
		mcp.WithDescription("List role resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// RoleHandler implements the handler for the Role tool
func RoleHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeRole,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllRole(ctx, "")

		},
	)
}
