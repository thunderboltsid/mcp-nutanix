package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// UserGroup defines the UserGroup tool
func UserGroup() mcp.Tool {
	return mcp.NewTool("usergroups",
		mcp.WithDescription("List usergroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// UserGroupHandler implements the handler for the UserGroup tool
func UserGroupHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeUserGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllUserGroup(ctx, "")

		},
	)
}
