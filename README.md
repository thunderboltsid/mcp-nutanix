# MCP Nutanix

A Model Context Protocol (MCP) server for interacting with Nutanix Prism Central APIs through Large Language Models (LLMs).

## ⚠️ Disclaimer

**THIS IS AN EXPERIMENTAL PROJECT**

This project was created as a personal project to explore the capabilities of the Model Context Protocol frameworks in Go. It is:

- **NOT** an official Nutanix product or tool
- **NOT** supported, endorsed, or maintained by Nutanix
- **NOT** suitable for production environments
- **PROVIDED AS-IS** with no warranties or guarantees

**USE AT YOUR OWN RISK**: The author takes no responsibility for any issues, damages, or outages that may result from using this code.

## Overview

This MCP server allows LLMs to interact with Nutanix Prism Central by:

1. Connecting to a Prism Central instance with user credentials
2. Listing various resources (VMs, Clusters, Hosts, etc.)
3. Retrieving specific resource details via URI-based access

The implementation uses the [Prism Go Client](https://github.com/nutanix-cloud-native/prism-go-client) to communicate with Prism Central and the [MCP Go library](https://github.com/mark3labs/mcp-go) to implement the Model Context Protocol.

## Getting Started

### Prerequisites

- Go 1.23 or higher
- Access to a Nutanix Prism Central instance
- Tools like `make` and `go fmt` for building

### Building

```bash
# Clone the repository
git clone https://github.com/thunderboltsid/mcp-nutanix.git
cd mcp-nutanix

# Build the MCP server
make build
```

### Running

```bash
./bin/mcp-nutanix
```

The server will start and prompt for Prism Central credentials.

## Usage

Once the MCP server is running and connected to your Prism Central instance, LLMs can interact with it through the MCP protocol.

### Resource Listing

To list resources, use the appropriate tool:

```
vms
clusters
hosts
images
subnets
```

The LLM will receive a JSON list of resources that it can parse and analyze.

### Resource Access

To access a specific resource, use a resource URI:

```
vm://{uuid}
cluster://{uuid}
host://{uuid}
```

The LLM will receive detailed JSON information about the specific resource.

## Development

### Project Structure

```
mcp-nutanix/
├── bin/                  # Compiled binaries
├── internal/             # Internal packages
│   ├── client/           # Prism Central client handling
│   ├── codegen/          # Code generation utilities
│   └── json/             # JSON helpers
├── pkg/                  # components
│   ├── prompts/          # MCP prompt implementations
│   ├── resources/        # Resource handlers
│   └── tools/            # Tool handlers
└── Makefile              # Build and utility commands
```

### Code Generation

The project uses code generation to create resource and tool handlers. To update these:

```bash
make generate
```

## Limitations

- Response size is limited by the MCP protocol
- Some resources with large response sizes may cause errors
- No pagination support in the current implementation
- Only supports read operations, no create/update/delete

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Nutanix](https://www.nutanix.com/) for creating the Prism API
- [Mark3Labs](https://github.com/mark3labs) for the MCP Go library
- [Nutanix Cloud Native](https://github.com/nutanix-cloud-native) for the Prism Go Client

## Contributing

This is an experimental project with no formal contribution process. Feel free to create issues or pull requests.