package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// User defines the User tool
func User() mcp.Tool {
	return mcp.NewTool("users",
		mcp.WithDescription("List user resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// UserHandler implements the handler for the User tool
func UserHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeUser,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllUser(ctx, "")

		},
	)
}
