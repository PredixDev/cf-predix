package main

import (
	"fmt"

	"github.com/cloudfoundry/cli/plugin"
)

type Predix struct{}

func (c *Predix) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Predix",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "predix",
				HelpText: "Sets the target API endpoint to a Predix PoP",
				UsageDetails: plugin.Usage{
					Usage: "predix - Prompts for options and sets the API endpoint\n    cf predix",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(Predix))
}

func (c *Predix) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "predix" {
		var choice int
		fmt.Println("1. US West")
		fmt.Println("2. US East")
		fmt.Println("3. Japan")
		fmt.Println("4. UK")
		fmt.Printf("Choose the PoP to set> ")
		fmt.Scanf("%d", &choice)

		var api string
		switch choice {
		case 1:
			api = "https://api.system.aws-usw02-pr.ice.predix.io"
		case 2:
			api = "https://api.system.asv-pr.ice.predix.io"
		case 3:
			api = "https://api.system.aws-jp01-pr.ice.predix.io"
		case 4:
			api = "https://api.system.dc-uk01-pr.ice.predix.io"
		default:
			fmt.Println("Invalid choice")
			return
		}

		cliConnection.CliCommand("login", "-a", api)
	}
}
