package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// AddressGroup defines the AddressGroup resource template
func AddressGroup() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeAddressGroup))+"{uuid}",
		string(ResourceTypeAddressGroup),
		mcp.WithTemplateDescription("Address Group resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// AddressGroupHandler implements the handler for the AddressGroup resource
func AddressGroupHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeAddressGroup, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the AddressGroup
		return client.V3().GetAddressGroup(ctx, uuid)
	})
}
