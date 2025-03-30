package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Subnet defines the Subnet tool
func Subnet() mcp.Tool {
	return mcp.NewTool("subnets",
		mcp.WithDescription("List subnet resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// SubnetHandler implements the handler for the Subnet tool
func SubnetHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeSubnet,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Special case for Subnet which has an extra parameter
			return client.V3().ListAllSubnet(ctx, "", nil)

		},
	)
}
