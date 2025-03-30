package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Role defines the Role resource template
func Role() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeRole))+"{uuid}",
		string(ResourceTypeRole),
		mcp.WithTemplateDescription("Role resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// RoleHandler implements the handler for the Role resource
func RoleHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeRole, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the Role
		return client.V3().GetRole(ctx, uuid)
	})
}
