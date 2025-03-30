package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	v3 "github.com/nutanix-cloud-native/prism-go-client/v3"
)

// Category defines the Category tool
func Category() mcp.Tool {
	return mcp.NewTool("categorys",
		mcp.WithDescription("List category resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// CategoryHandler implements the handler for the Category tool
func CategoryHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeCategory,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Special case for Category which takes CategoryListMetadata
			var length int64 = 100
			metadata := &v3.CategoryListMetadata{
				Length: &length,
			}
			return client.V3().ListCategories(ctx, metadata)

		},
	)
}
