package resources

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/thunderboltsid/mcp-nutanix/pkg/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func Cluster() mcp.ResourceTemplate {
	return mcp.NewResourceTemplate(
		"cluster://{uuid}",
		"cluster",
		mcp.WithTemplateDescription("Prism Element (AOS) Cluster resource"),
		mcp.WithTemplateMIMEType("application/json"),
	)
}

func ClusterHandler() server.ResourceTemplateHandlerFunc {
	return func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		clusterUUID := extractIDFromURI(request.Params.URI)
		if clusterUUID == "" {
			return nil, fmt.Errorf("URI must contain a UUID")
		}

		resp, err := client.GetPrismClient().V3().GetCluster(ctx, clusterUUID)
		if err != nil {
			return nil, err
		}

		enc, err := json.Marshal(resp)
		if err != nil {
			return nil, fmt.Errorf("failed to encode cluster response: %w", err)
		}

		return []mcp.ResourceContents{
			&mcp.TextResourceContents{
				URI:      request.Params.URI,
				MIMEType: "application/json",
				Text:     string(enc),
			},
		}, nil
	}
}
