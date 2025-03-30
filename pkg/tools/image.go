package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Image defines the Image tool
func Image() mcp.Tool {
	return mcp.NewTool("images",
		mcp.WithDescription("List image resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ImageHandler implements the handler for the Image tool
func ImageHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeImage,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllImage(ctx, "")

		},
	)
}
