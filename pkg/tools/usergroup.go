package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// UserGroup defines the UserGroup tool
func UserGroupList() mcp.Tool {
	return mcp.NewTool("usergroup_list",
		mcp.WithDescription("List usergroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// UserGroupListHandler implements the handler for the UserGroup list tool
func UserGroupListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeUserGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllUserGroup(ctx, "")

		},
	)
}

// UserGroupCount defines the UserGroup count tool
func UserGroupCount() mcp.Tool {
	return mcp.NewTool("usergroup_count",
		mcp.WithDescription("Count usergroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// UserGroupCountHandler implements the handler for the UserGroup count tool
func UserGroupCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeUserGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllUserGroup(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "UserGroup",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
