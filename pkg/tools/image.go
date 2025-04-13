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
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllImage(ctx, "")

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
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllImage(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "Image",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
