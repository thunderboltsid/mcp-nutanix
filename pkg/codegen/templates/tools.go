package templates

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Templates for tool implementations
const toolTemplate = `package tools

import (
    "context"
    "fmt"

    "github.com/thunderboltsid/mcp-nutanix/internal/client"
    "github.com/thunderboltsid/mcp-nutanix/pkg/resources"

    "github.com/mark3labs/mcp-go/mcp"
    "github.com/mark3labs/mcp-go/server"
    "github.com/nutanix-cloud-native/prism-go-client/v3"
)

// {{.Name}} defines the {{.Name}} tool
func {{.Name}}() mcp.Tool {
    return mcp.NewTool("{{.ResourceType}}s",
        mcp.WithDescription("List {{.ResourceType}} resources"),
        mcp.WithString("filter",
            mcp.Description("Optional text filter (interpreted by LLM)"),
        ),
    )
}

// {{.Name}}Handler implements the handler for the {{.Name}} tool
func {{.Name}}Handler() server.ToolHandlerFunc {
    return CreateToolHandler(
        resources.ResourceType{{.Name}},
        // Define the ListResourceFunc implementation
        func(ctx context.Context, client *client.NutanixClient, filter string) (interface{}, error) {
            {{if eq .Name "Host"}}
            // Special case for Host which doesn't take a filter
            return client.V3().{{.ClientListAllFunc}}(ctx)
            {{else if eq .Name "Subnet"}}
            // Special case for Subnet which has an extra parameter
            return client.V3().{{.ClientListAllFunc}}(ctx, "", nil)
            {{else if eq .Name "Category"}}
            // Special case for Category which takes CategoryListMetadata
            var length int64 = 100
            metadata := &v3.CategoryListMetadata{
                Length: &length,
            }
            return client.V3().{{.ClientListFunc}}(ctx, metadata)
            {{else if .HasListAllFunc}}
            // Use ListAll function to get all resources
            return client.V3().{{.ClientListAllFunc}}(ctx, "")
            {{else}}
            // Create DSMetadata without filter
            var length int64 = 100
            metadata := &v3.DSMetadata{
                Length: &length,
            }
            
            return client.V3().{{.ClientListFunc}}(ctx, metadata)
            {{end}}
        },
    )
}
`

// GenerateToolFiles generates tool files for all Nutanix resources that support listing
func GenerateToolFiles(baseDir string) error {
	resources := GetResourceDefinitions()

	// Create the tools directory if it doesn't exist
	toolsDir := fmt.Sprintf("%s/pkg/tools", baseDir)
	err := os.MkdirAll(toolsDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating tools directory: %w", err)
	}

	// Parse the tool template
	toolTmpl, err := template.New("tool").Parse(toolTemplate)
	if err != nil {
		return fmt.Errorf("error parsing tool template: %w", err)
	}

	// Generate tool files
	for _, res := range resources {
		// Skip resources that don't support listing
		if !res.HasListFunc && !res.HasListAllFunc {
			fmt.Printf("Skipping tool generation for %s: no list capability\n", res.Name)
			continue
		}

		// Create tool file
		toolFilePath := fmt.Sprintf("%s/%s.go", toolsDir, strings.ToLower(res.Name))
		toolFile, err := os.Create(toolFilePath)
		if err != nil {
			fmt.Printf("Error creating tool file for %s: %v\n", res.Name, err)
			continue
		}
		defer toolFile.Close()

		// Execute the template
		err = toolTmpl.Execute(toolFile, res)
		if err != nil {
			fmt.Printf("Error executing tool template for %s: %v\n", res.Name, err)
		}
	}

	return nil
}
