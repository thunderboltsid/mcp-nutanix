package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ProtectionRule defines the ProtectionRule resource template
func ProtectionRule() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeProtectionRule))+"{uuid}",
		string(ResourceTypeProtectionRule),
		mcp.WithTemplateDescription("Protection Rule resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// ProtectionRuleHandler implements the handler for the ProtectionRule resource
func ProtectionRuleHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeProtectionRule, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the ProtectionRule
		return client.V3().GetProtectionRule(ctx, uuid)
	})
}
