package main

import (
	"fmt"
	"os"

	"github.com/PredixDev/cf-predix/version"
	"github.com/cloudfoundry/cli/plugin"
)

type Predix struct{}

func (c *Predix) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:    "Predix",
		Version: version.GetVersion(),
		Commands: []plugin.Command{
			{
				Name:     "predix",
				HelpText: "Sets the target API endpoint to a Predix PoP",
				UsageDetails: plugin.Usage{
					Usage: "predix - Prompts for options and sets the API endpoint\n    cf predix",
				},
			},
			{
				Name:     "version",
				HelpText: "Prints the version info",
				UsageDetails: plugin.Usage{
					Usage: "predix - version",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(Predix))
}

func (c *Predix) Run(cliConnection plugin.CliConnection, args []string) {
	switch args[0] {
	case "version":
		fmt.Printf("Predix cf plugin. Version %s [Commit: %s]", version.VersionString, version.CommitSha)
		break
	case "predix":
		choosePop(cliConnection)
		break
	default:
		fmt.Printf("Not sure how to run the command %s\n", args[0])
		os.Exit(1)
		return
	}
}
