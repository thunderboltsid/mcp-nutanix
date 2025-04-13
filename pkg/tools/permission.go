package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Permission defines the Permission tool
func PermissionList() mcp.Tool {
	return mcp.NewTool("permission_list",
		mcp.WithDescription("List permission resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// PermissionListHandler implements the handler for the Permission list tool
func PermissionListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypePermission,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllPermission(ctx, "")

		},
	)
}

// PermissionCount defines the Permission count tool
func PermissionCount() mcp.Tool {
	return mcp.NewTool("permission_count",
		mcp.WithDescription("Count permission resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// PermissionCountHandler implements the handler for the Permission count tool
func PermissionCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypePermission,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllPermission(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "Permission",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
