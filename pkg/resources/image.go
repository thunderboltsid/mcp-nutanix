package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

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
		// Use converged API to get Image by ID
		return client.Converged().Images.Get(ctx, uuid)
	})
}
