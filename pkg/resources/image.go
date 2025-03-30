package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Image defines the Image resource template
func Image() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeImage))+"{uuid}",
		string(ResourceTypeImage),
		mcp.WithTemplateDescription("Image resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// ImageHandler implements the handler for the Image resource
func ImageHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeImage, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the Image
		return client.V3().GetImage(ctx, uuid)
	})
}
