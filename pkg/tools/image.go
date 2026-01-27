package tools

import (
	"context"
	"reflect"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Image defines the Image tool
func ImageList() mcp.Tool {
	return mcp.NewTool("image_list",
		mcp.WithDescription("List image resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ImageListHandler implements the handler for the Image list tool
func ImageListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeImage,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			var allImages []interface{}
			page := 0
			limit := 100

			for {
				resp, err := client.V4().ImagesApiInstance.ListImages(&page, &limit, nil, nil, nil)
				if err != nil {
					return nil, err
				}

				if resp.Data != nil {
					if dataValue := resp.Data.GetValue(); dataValue != nil {
						v := reflect.ValueOf(dataValue)
						if v.Kind() == reflect.Slice {
							for i := 0; i < v.Len(); i++ {
								allImages = append(allImages, v.Index(i).Interface())
							}
						}
					}
				}

				if resp.Metadata != nil && resp.Metadata.Links != nil {
					hasNext := false
					for _, link := range resp.Metadata.Links {
						if link.Rel != nil && *link.Rel == "next" {
							hasNext = true
							break
						}
					}
					if !hasNext {
						break
					}
					page++
				} else {
					break
				}
			}

			return map[string]interface{}{
				"images": allImages,
				"count":  len(allImages),
			}, nil
		},
	)
}

// ImageCount defines the Image count tool
func ImageCount() mcp.Tool {
	return mcp.NewTool("image_count",
		mcp.WithDescription("Count image resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ImageCountHandler implements the handler for the Image count tool
func ImageCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeImage,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			totalCount := 0
			page := 0
			limit := 100

			for {
				resp, err := client.V4().ImagesApiInstance.ListImages(&page, &limit, nil, nil, nil)
				if err != nil {
					return nil, err
				}

				if resp.Data != nil {
					if dataValue := resp.Data.GetValue(); dataValue != nil {
						v := reflect.ValueOf(dataValue)
						if v.Kind() == reflect.Slice {
							totalCount += v.Len()
						}
					}
				}

				if resp.Metadata != nil && resp.Metadata.Links != nil {
					hasNext := false
					for _, link := range resp.Metadata.Links {
						if link.Rel != nil && *link.Rel == "next" {
							hasNext = true
							break
						}
					}
					if !hasNext {
						break
					}
					page++
				} else {
					break
				}
			}

			return map[string]interface{}{
				"resource_type": "Image",
				"count":         totalCount,
			}, nil
		},
	)
}
