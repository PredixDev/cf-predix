package main

import (
	"fmt"
	"os"

	"github.com/PredixDev/cf-predix/pop"
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
		},
	}
}

func main() {
	plugin.Start(new(Predix))
}

func (c *Predix) Run(cliConnection plugin.CliConnection, args []string) {
	switch args[0] {
	case "predix":
		pop.ChoosePop(cliConnection)
		break
	default:
		fmt.Printf("Not sure how to run the command %q\n", args[0])
		os.Exit(1)
		return
	}
}
