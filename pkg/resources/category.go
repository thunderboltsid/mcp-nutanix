package resources

import (
	"context"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Category defines the Category resource template
func Category() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		string(ResourceURIPrefix(ResourceTypeCategory))+"{uuid}",
		string(ResourceTypeCategory),
		mcp.WithTemplateDescription("Category resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

// CategoryHandler implements the handler for the Category resource
func CategoryHandler() server.ResourceTemplateHandlerFunc {
	return CreateResourceHandler(ResourceTypeCategory, func(ctx context.Context, client *client.NutanixClient, uuid string) (interface{}, error) {
		// Get the Category
		return client.V3().GetCategoryKey(ctx, uuid)
	})
}
