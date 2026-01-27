package tools

import (
	"context"
	"reflect"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Host defines the Host tool
func HostList() mcp.Tool {
	return mcp.NewTool("host_list",
		mcp.WithDescription("List host resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// HostListHandler implements the handler for the Host list tool
func HostListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeHost,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			var allHosts []interface{}
			page := 0
			limit := 100

			for {
				resp, err := client.V4().ClustersApiInstance.ListHosts(&page, &limit, nil, nil, nil, nil)
				if err != nil {
					return nil, err
				}

				if resp.Data != nil {
					if dataValue := resp.Data.GetValue(); dataValue != nil {
						v := reflect.ValueOf(dataValue)
						if v.Kind() == reflect.Slice {
							for i := 0; i < v.Len(); i++ {
								allHosts = append(allHosts, v.Index(i).Interface())
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
				"hosts": allHosts,
				"count": len(allHosts),
			}, nil
		},
	)
}

// HostCount defines the Host count tool
func HostCount() mcp.Tool {
	return mcp.NewTool("host_count",
		mcp.WithDescription("Count host resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// HostCountHandler implements the handler for the Host count tool
func HostCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeHost,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			totalCount := 0
			page := 0
			limit := 100

			for {
				resp, err := client.V4().ClustersApiInstance.ListHosts(&page, &limit, nil, nil, nil, nil)
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
				"resource_type": "Host",
				"count":         totalCount,
			}, nil
		},
	)
}
