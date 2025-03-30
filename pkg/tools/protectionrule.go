package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ProtectionRule defines the ProtectionRule tool
func ProtectionRule() mcp.Tool {
	return mcp.NewTool("protectionrules",
		mcp.WithDescription("List protectionrule resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ProtectionRuleHandler implements the handler for the ProtectionRule tool
func ProtectionRuleHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeProtectionRule,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllProtectionRules(ctx, "")

		},
	)
}
