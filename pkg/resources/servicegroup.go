package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// ServiceGroup defines the ServiceGroup resource template
func ServiceGroup() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeServiceGroup))+"{uuid}",
		string(ResourceTypeServiceGroup),
		mcp.WithTemplateDescription("Service Group resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// ServiceGroupHandler implements the handler for the ServiceGroup resource
func ServiceGroupHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeServiceGroup, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the ServiceGroup
		return client.V3().GetServiceGroup(ctx, uuid)
	})
}
