package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// UserGroup defines the UserGroup resource template
func UserGroup() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeUserGroup))+"{uuid}",
		string(ResourceTypeUserGroup),
		mcp.WithTemplateDescription("User Group resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// UserGroupHandler implements the handler for the UserGroup resource
func UserGroupHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeUserGroup, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the UserGroup
		return client.V3().GetUserGroup(ctx, uuid)
	})
}
