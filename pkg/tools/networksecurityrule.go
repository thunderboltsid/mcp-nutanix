package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NetworkSecurityRule defines the NetworkSecurityRule tool
func NetworkSecurityRule() mcp.Tool {
	return mcp.NewTool("networksecurityrules",
		mcp.WithDescription("List networksecurityrule resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// NetworkSecurityRuleHandler implements the handler for the NetworkSecurityRule tool
func NetworkSecurityRuleHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeNetworkSecurityRule,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllNetworkSecurityRule(ctx, "")

		},
	)
}
