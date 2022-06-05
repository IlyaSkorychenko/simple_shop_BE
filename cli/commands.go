package cli

type cliCommand struct {
	flagValue string
	flags     []string
	action    func(flagValue string)
}

func (c cliCommand) runAction() {
	c.action(c.flagValue)
}

func (c cliCommand) validateFlags() bool {
	for _, availableFlag := range c.flags {
		if c.flagValue == availableFlag {
			return true
		}
	}

	return false
}

func migrationCommand(flagValue string) cliCommand {
	return cliCommand{
		flagValue: flagValue,
		flags:     []string{up, down},
		action:    runMigration,
	}
}

func seederCommand(flagValue string) cliCommand {
	return cliCommand{
		flagValue: flagValue,
		flags:     []string{run},
		action:    runSeeder,
	}
}
