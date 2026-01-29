package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// CategoryList defines the Category list tool
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
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			categories, err := client.Converged().Categories.List(ctx)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"categories": categories,
				"count":      len(categories),
			}, nil
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
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			categories, err := client.Converged().Categories.List(ctx)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"resource_type": "Category",
				"count":         len(categories),
			}, nil
		},
	)
}
