package prompts

import (
	"context"
	"fmt"

	"github.com/thunderboltsid/mcp-nutanix/internal/client"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func SetCredentials() mcp.Prompt {
	return mcp.NewPrompt("credentials",
		mcp.WithPromptDescription("Credentials for Prism Central"),
		mcp.WithArgument("endpoint",
			mcp.RequiredArgument(),
			mcp.ArgumentDescription("Prism Central endpoint"),
		),
		mcp.WithArgument("username",
			mcp.RequiredArgument(),
			mcp.ArgumentDescription("Username of the Prism Central user"),
		),
		mcp.WithArgument("password",
			mcp.RequiredArgument(),
			mcp.ArgumentDescription("Password of the Prism Central user"),
		),
		mcp.WithArgument("insecure",
			mcp.ArgumentDescription("Skip TLS verification (true/false)"),
		),
	)
}

func SetCredentialsResponse() server.PromptHandlerFunc {
	return func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		endpoint := request.Params.Arguments["endpoint"]
		username := request.Params.Arguments["username"]
		password := request.Params.Arguments["password"]
		insecure := request.Params.Arguments["insecure"]

		client.PrismClientProvider.UpdateValue("endpoint", endpoint)
		client.PrismClientProvider.UpdateValue("username", username)
		client.PrismClientProvider.UpdateValue("password", password)
		client.PrismClientProvider.UpdateValue("insecure", insecure)

		client.Init(client.PrismClientProvider)

		// Validate the credentials
		pcInfo, err := client.GetPrismClient().V3().GetPrismCentral(ctx)
		if err != nil {
			return mcp.NewGetPromptResult(
				"Failed to connect to Prism Central",
				[]mcp.PromptMessage{
					mcp.NewPromptMessage(
						mcp.RoleAssistant,
						mcp.NewTextContent(fmt.Sprintf("Failed to connect to Prism Central: %s. Please check your credentials and try again.", err.Error())),
					),
				},
			), nil
		}

		return mcp.NewGetPromptResult(
			"Connected to Prism Central",
			[]mcp.PromptMessage{
				mcp.NewPromptMessage(
					mcp.RoleAssistant,
					mcp.NewTextContent(fmt.Sprintf("Successfully connected to Prism Central %s at %s. You can now use the tools to interact with your Nutanix environment.", pcInfo.Resources.ClusterUUID, endpoint)),
				),
			},
		), nil
	}
}
