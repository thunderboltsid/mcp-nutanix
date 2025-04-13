package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ProtectionRule defines the ProtectionRule tool
func ProtectionRuleList() mcp.Tool {
	return mcp.NewTool("protectionrule_list",
		mcp.WithDescription("List protectionrule resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ProtectionRuleListHandler implements the handler for the ProtectionRule list tool
func ProtectionRuleListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeProtectionRule,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllProtectionRules(ctx, "")

		},
	)
}

// ProtectionRuleCount defines the ProtectionRule count tool
func ProtectionRuleCount() mcp.Tool {
	return mcp.NewTool("protectionrule_count",
		mcp.WithDescription("Count protectionrule resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ProtectionRuleCountHandler implements the handler for the ProtectionRule count tool
func ProtectionRuleCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeProtectionRule,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllProtectionRules(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "ProtectionRule",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
