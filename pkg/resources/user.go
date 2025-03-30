package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// User defines the User resource template
func User() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeUser))+"{uuid}",
		string(ResourceTypeUser),
		mcp.WithTemplateDescription("User resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// UserHandler implements the handler for the User resource
func UserHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeUser, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the User
		return client.V3().GetUser(ctx, uuid)
	})
}
