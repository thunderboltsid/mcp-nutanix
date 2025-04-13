package tools

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NetworkSecurityRule defines the NetworkSecurityRule tool
func NetworkSecurityRuleList() mcp.Tool {
	return mcp.NewTool("networksecurityrule_list",
		mcp.WithDescription("List networksecurityrule resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// NetworkSecurityRuleListHandler implements the handler for the NetworkSecurityRule list tool
func NetworkSecurityRuleListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeNetworkSecurityRule,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			return client.V3().ListAllNetworkSecurityRule(ctx, "")

		},
	)
}

// NetworkSecurityRuleCount defines the NetworkSecurityRule count tool
func NetworkSecurityRuleCount() mcp.Tool {
	return mcp.NewTool("networksecurityrule_count",
		mcp.WithDescription("Count networksecurityrule resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// NetworkSecurityRuleCountHandler implements the handler for the NetworkSecurityRule count tool
func NetworkSecurityRuleCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeNetworkSecurityRule,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Use ListAll function to get all resources
			resp, err := client.V3().ListAllNetworkSecurityRule(ctx, "")

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "NetworkSecurityRule",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
