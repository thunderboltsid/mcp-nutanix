package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	v3 "github.com/nutanix-cloud-native/prism-go-client/v3"
)

// Category defines the Category tool
func CategoryList() mcp.Tool {
	return mcp.NewTool("category_list",
		mcp.WithDescription("List category resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// CategoryListHandler implements the handler for the Category list tool
func CategoryListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeCategory,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Special case for Category which takes CategoryListMetadata
			metadata := &v3.CategoryListMetadata{}
			return client.V3().ListCategories(ctx, metadata)

		},
	)
}

// CategoryCount defines the Category count tool
func CategoryCount() mcp.Tool {
	return mcp.NewTool("category_count",
		mcp.WithDescription("Count category resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// CategoryCountHandler implements the handler for the Category count tool
func CategoryCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeCategory,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Special case for Category which takes CategoryListMetadata
			metadata := &v3.CategoryListMetadata{}
			resp, err := client.V3().ListCategories(ctx, metadata)

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "Category",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
