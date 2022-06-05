package cli

import (
	"flag"
	"fmt"
)

type commandsNode struct {
	command   func(flagValue string) cliCommand
	flagValue string
}

var commandNodes = map[string]*commandsNode{
	migrationFlag: {migrationCommand, ""},
	seederFlag:    {seederCommand, ""},
}

const (
	defFlagValue  = "unset"
	migrationFlag = "migration"
	seederFlag    = "seeder"
)

func init() {
	flag.StringVar(&commandNodes[migrationFlag].flagValue, migrationFlag, defFlagValue, "Run migration")
	flag.StringVar(&commandNodes[seederFlag].flagValue, seederFlag, defFlagValue, "Run seeder")
}

func Run() {
	for flagName, node := range commandNodes {
		command := node.command(node.flagValue)

		if node.flagValue == defFlagValue {
			continue
		}

		if !command.validateFlags() {
			panic(fmt.Sprintf("%s: unexpected value", flagName))
		}

		command.runAction()

		return
	}

	fmt.Println("!! No commands !!")
}
