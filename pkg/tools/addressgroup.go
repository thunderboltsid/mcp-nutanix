package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AddressGroup defines the AddressGroup tool
func AddressGroup() mcp.Tool {
	return mcp.NewTool("addressgroups",
		mcp.WithDescription("List addressgroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// AddressGroupHandler implements the handler for the AddressGroup tool
func AddressGroupHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeAddressGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllAddressGroups(ctx, "")

		},
	)
}
