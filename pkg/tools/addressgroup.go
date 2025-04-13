package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AddressGroup defines the AddressGroup tool
func AddressGroupList() mcp.Tool {
	return mcp.NewTool("addressgroup_list",
		mcp.WithDescription("List addressgroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// AddressGroupListHandler implements the handler for the AddressGroup list tool
func AddressGroupListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeAddressGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllAddressGroups(ctx, "")

		},
	)
}

// AddressGroupCount defines the AddressGroup count tool
func AddressGroupCount() mcp.Tool {
	return mcp.NewTool("addressgroup_count",
		mcp.WithDescription("Count addressgroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// AddressGroupCountHandler implements the handler for the AddressGroup count tool
func AddressGroupCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeAddressGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllAddressGroups(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "AddressGroup",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
