package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// User defines the User tool
func UserList() mcp.Tool {
	return mcp.NewTool("user_list",
		mcp.WithDescription("List user resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// UserListHandler implements the handler for the User list tool
func UserListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeUser,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllUser(ctx, "")

		},
	)
}

// UserCount defines the User count tool
func UserCount() mcp.Tool {
	return mcp.NewTool("user_count",
		mcp.WithDescription("Count user resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// UserCountHandler implements the handler for the User count tool
func UserCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeUser,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllUser(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "User",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
