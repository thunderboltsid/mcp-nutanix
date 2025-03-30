package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ServiceGroup defines the ServiceGroup tool
func ServiceGroup() mcp.Tool {
	return mcp.NewTool("servicegroups",
		mcp.WithDescription("List servicegroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ServiceGroupHandler implements the handler for the ServiceGroup tool
func ServiceGroupHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeServiceGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllServiceGroups(ctx, "")

		},
	)
}
