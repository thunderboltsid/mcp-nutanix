package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Project defines the Project tool
func Project() mcp.Tool {
	return mcp.NewTool("projects",
		mcp.WithDescription("List project resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ProjectHandler implements the handler for the Project tool
func ProjectHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeProject,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllProject(ctx, "")

		},
	)
}
