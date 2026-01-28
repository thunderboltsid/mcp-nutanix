package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Image defines the Image tool
func ImageList() mcp.Tool {
	return mcp.NewTool("image_list",
		mcp.WithDescription("List image resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ImageListHandler implements the handler for the Image list tool
func ImageListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeImage,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			images, err := client.Converged().Images.List(ctx)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"images": images,
				"count":  len(images),
			}, nil
		},
	)
}

// ImageCount defines the Image count tool
func ImageCount() mcp.Tool {
	return mcp.NewTool("image_count",
		mcp.WithDescription("Count image resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ImageCountHandler implements the handler for the Image count tool
func ImageCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeImage,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			images, err := client.Converged().Images.List(ctx)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{
				"resource_type": "Image",
				"count":         len(images),
			}, nil
		},
	)
}
