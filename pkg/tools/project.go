package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Project defines the Project tool
func ProjectList() mcp.Tool {
	return mcp.NewTool("project_list",
		mcp.WithDescription("List project resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ProjectListHandler implements the handler for the Project list tool
func ProjectListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeProject,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllProject(ctx, "")

		},
	)
}

// ProjectCount defines the Project count tool
func ProjectCount() mcp.Tool {
	return mcp.NewTool("project_count",
		mcp.WithDescription("Count project resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ProjectCountHandler implements the handler for the Project count tool
func ProjectCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeProject,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllProject(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "Project",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
