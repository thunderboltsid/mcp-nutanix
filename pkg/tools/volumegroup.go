package tools

import (
	"context"

	v3 "github.com/nutanix-cloud-native/prism-go-client/v3"
	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// VolumeGroup defines the VolumeGroup tool
func VolumeGroup() mcp.Tool {
	return mcp.NewTool("volumegroups",
		mcp.WithDescription("List volumegroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// VolumeGroupHandler implements the handler for the VolumeGroup tool
func VolumeGroupHandler() server.ToolHandlerFunc {
	return CreateToolHandler(
		resources.ResourceTypeVolumeGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Create DSMetadata without filter
			var length int64 = 100
			metadata := &v3.DSMetadata{
				Length: &length,
			}

			return client.V3().ListVolumeGroup(ctx, metadata)

		},
	)
}
