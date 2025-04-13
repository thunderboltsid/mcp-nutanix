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
func VolumeGroupList() mcp.Tool {
	return mcp.NewTool("volumegroup_list",
		mcp.WithDescription("List volumegroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// VolumeGroupListHandler implements the handler for the VolumeGroup list tool
func VolumeGroupListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeVolumeGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Create DSMetadata without filter
			metadata := &v3.DSMetadata{}

			return client.V3().ListVolumeGroup(ctx, metadata)

		},
	)
}

// VolumeGroupCount defines the VolumeGroup count tool
func VolumeGroupCount() mcp.Tool {
	return mcp.NewTool("volumegroup_count",
		mcp.WithDescription("Count volumegroup resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// VolumeGroupCountHandler implements the handler for the VolumeGroup count tool
func VolumeGroupCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeVolumeGroup,
		// Define the ListResourceFunc implementation
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {

			// Create DSMetadata without filter
			metadata := &v3.DSMetadata{}

			resp, err := client.V3().ListVolumeGroup(ctx, metadata)

			if err != nil {
				return nil, err
			}

			res := map[string]interface{}{
				"resource_type": "VolumeGroup",
				"count":         len(resp.Entities),
				"metadata":      resp.Metadata,
			}

			return res, nil
		},
	)
}
