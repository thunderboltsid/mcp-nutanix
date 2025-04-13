package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Subnet defines the Subnet tool
func SubnetList() mcp.Tool {
	return mcp.NewTool("subnet_list",
		mcp.WithDescription("List subnet resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// SubnetListHandler implements the handler for the Subnet list tool
func SubnetListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeSubnet,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Special case for Subnet which has an extra parameter
			return client.V3().ListAllSubnet(ctx, "", nil)

		},
	)
}

// SubnetCount defines the Subnet count tool
func SubnetCount() mcp.Tool {
	return mcp.NewTool("subnet_count",
		mcp.WithDescription("Count subnet resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// SubnetCountHandler implements the handler for the Subnet count tool
func SubnetCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeSubnet,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Special case for Subnet which has an extra parameter
			resp, err := client.V3().ListAllSubnet(ctx, "", nil)

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "Subnet",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
