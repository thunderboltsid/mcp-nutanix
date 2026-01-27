package tools

import (
	"context"
	"reflect"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"
	"github.com/thunderboltsid/mcp-nutanix/pkg/resources"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Cluster defines the Cluster tool
func ClusterList() mcp.Tool {
	return mcp.NewTool("cluster_list",
		mcp.WithDescription("List cluster resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ClusterListHandler implements the handler for the Cluster list tool
func ClusterListHandler() server.ToolHandlerFunc {
	return CreateListToolHandler(
		resources.ResourceTypeCluster,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			var allClusters []interface{}
			page := 0
			limit := 100

			for {
				resp, err := client.V4().ClustersApiInstance.ListClusters(&page, &limit, nil, nil, nil, nil, nil)
				if err != nil {
					return nil, err
				}

				if resp.Data != nil {
					if dataValue := resp.Data.GetValue(); dataValue != nil {
						v := reflect.ValueOf(dataValue)
						if v.Kind() == reflect.Slice {
							for i := 0; i < v.Len(); i++ {
								allClusters = append(allClusters, v.Index(i).Interface())
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
				"clusters": allClusters,
				"count":    len(allClusters),
			}, nil
		},
	)
}

// ClusterCount defines the Cluster count tool
func ClusterCount() mcp.Tool {
	return mcp.NewTool("cluster_count",
		mcp.WithDescription("Count cluster resources"),
		mcp.WithString("filter",
			mcp.Description("Optional text filter (interpreted by LLM)"),
		),
	)
}

// ClusterCountHandler implements the handler for the Cluster count tool
func ClusterCountHandler() server.ToolHandlerFunc {
	return CreateCountToolHandler(
		resources.ResourceTypeCluster,
		func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
			totalCount := 0
			page := 0
			limit := 100

			for {
				resp, err := client.V4().ClustersApiInstance.ListClusters(&page, &limit, nil, nil, nil, nil, nil)
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
				"resource_type": "Cluster",
				"count":         totalCount,
			}, nil
		},
	)
}
