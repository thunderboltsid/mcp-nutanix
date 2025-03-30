package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// NetworkSecurityRule defines the NetworkSecurityRule resource template
func NetworkSecurityRule() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeNetworkSecurityRule))+"{uuid}",
		string(ResourceTypeNetworkSecurityRule),
		mcp.WithTemplateDescription("Network Security Rule resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// NetworkSecurityRuleHandler implements the handler for the NetworkSecurityRule resource
func NetworkSecurityRuleHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeNetworkSecurityRule, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the NetworkSecurityRule
		return client.V3().GetNetworkSecurityRule(ctx, uuid)
	})
}
