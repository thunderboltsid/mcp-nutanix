package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Project defines the Project resource template
func Project() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeProject))+"{uuid}",
		string(ResourceTypeProject),
		mcp.WithTemplateDescription("Project resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// ProjectHandler implements the handler for the Project resource
func ProjectHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeProject, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the Project
		return client.V3().GetProject(ctx, uuid)
	})
}
