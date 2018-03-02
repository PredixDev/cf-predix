package pop

import (
	"fmt"
	"log"
	"os"
)

type cliplugins interface {
	CliCommand(args ...string) ([]string, error)
}

func ChoosePop(cliConnection cliplugins) {
	var choice int
	for i, pop := range availablePops {
		fmt.Printf("%d - %s\n", i, pop.Label)
	}
	fmt.Printf("Choose the PoP to set> ")
	fmt.Scanf("%d", &choice)
	if choice > len(availablePops) {
		log.Printf("%d is an invalid choice\n", choice)
		os.Exit(1)
	}

	api := availablePops[choice].ApiEndpoint

	cliConnection.CliCommand("login", "-a", api)
}
