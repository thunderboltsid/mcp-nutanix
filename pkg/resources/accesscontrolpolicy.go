package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AccessControlPolicy defines the AccessControlPolicy resource template
func AccessControlPolicy() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeAccessControlPolicy))+"{uuid}",
		string(ResourceTypeAccessControlPolicy),
		mcp.WithTemplateDescription("Access Control Policy resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// AccessControlPolicyHandler implements the handler for the AccessControlPolicy resource
func AccessControlPolicyHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeAccessControlPolicy, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the AccessControlPolicy
		return client.V3().GetAccessControlPolicy(ctx, uuid)
	})
}
