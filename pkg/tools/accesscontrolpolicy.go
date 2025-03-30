package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AccessControlPolicy defines the AccessControlPolicy tool
func AccessControlPolicy() mcp.Tool {
	return mcp.NewTool("accesscontrolpolicys",
		mcp.WithDescription("List accesscontrolpolicy resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// AccessControlPolicyHandler implements the handler for the AccessControlPolicy tool
func AccessControlPolicyHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeAccessControlPolicy,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllAccessControlPolicy(ctx, "")

		},
	)
}
