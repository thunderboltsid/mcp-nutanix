package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ServiceGroup defines the ServiceGroup tool
func ServiceGroupList() mcp.Tool {
	return mcp.NewTool("servicegroup_list",
		mcp.WithDescription("List servicegroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ServiceGroupListHandler implements the handler for the ServiceGroup list tool
func ServiceGroupListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeServiceGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllServiceGroups(ctx, "")

		},
	)
}

// ServiceGroupCount defines the ServiceGroup count tool
func ServiceGroupCount() mcp.Tool {
	return mcp.NewTool("servicegroup_count",
		mcp.WithDescription("Count servicegroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ServiceGroupCountHandler implements the handler for the ServiceGroup count tool
func ServiceGroupCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeServiceGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllServiceGroups(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "ServiceGroup",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
