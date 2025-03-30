# Nutanix Prism Central MCP Code Generator

This package contains utilities for automatically generating Go code for the Nutanix Prism Central MCP server. It helps maintain consistency across resource handlers and reduces boilerplate code duplication.

## Purpose

The code generator creates:

1. Resource handlers for all Nutanix Prism Central resource types
2. Tool handlers for listing and retrieving these resources
3. Consistent interface implementations across all resources

## Usage

### Command Line Tool

To generate all resource and tool files:

```bash
go run cmd/codegen/main.go --output /path/to/project
```

Options:
- `--output`: The root directory of your project (default: current directory)

### Library Usage

You can also use the code generator programmatically in your own tools:

```go
import "github.com/thunderboltsid/mcp-nutanix/pkg/codegen"

func main() {
    // Generate files in the current directory
    err := codegen.GenerateResourceFiles(".")
    if err != nil {
        panic(err)
    }
}
```

## Resource Structure

Each generated resource consists of:

1. A resource handler in `pkg/resources` that implements `server.ResourceTemplateHandlerFunc`
2. A tool handler in `pkg/tools` that implements `server.ToolHandlerFunc`

The handlers use the generic handler functions from their respective packages to provide consistent behavior.

## Adding New Resources

To add a new resource type to the generator:

1. Edit `pkg/codegen/templates/codegen.go`
2. Add a new entry to the `GetResourceDefinitions()` function with:
    - Name: The resource name in Pascal case (e.g., "VolumeGroup")
    - ResourceType: The resource identifier in lowercase (e.g., "volumegroup")
    - Description: A brief description of the resource
    - ClientGetFunc: The V3 client method for retrieving a single resource
    - ClientListFunc: The V3 client method for listing resources
    - ClientListAllFunc: The V3 client method for listing all resources
    - FilterField: The field name used when filtering by name

3. Run the generator to create the new files

## Templates

The code generator uses Go templates to create the resource and tool handlers. These templates can be found in the `codegen.go` file and can be modified to change the structure of the generated code.