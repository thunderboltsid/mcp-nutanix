package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Permission defines the Permission resource template
func Permission() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypePermission))+"{uuid}",
		string(ResourceTypePermission),
		mcp.WithTemplateDescription("Permission resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// PermissionHandler implements the handler for the Permission resource
func PermissionHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypePermission, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the Permission
		return client.V3().GetPermission(ctx, uuid)
	})
}
