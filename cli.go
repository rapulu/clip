package clip

type Cli struct{
	version string
	rootCommand *Command
	errorHandler func(string, error) error
	preRunCommand func(*Cli) error
}


func(c *Cli) LongDescription(longdescription string) *Cli {
	c.rootCommand.longdescription = longdescription
	return c
}

func(c *Cli) Action(callback Action) *Cli {
	c.rootCommand.actionCallback = callback
	return c
}

func(c *Cli) Run(args ...string) error {
	if c.preRunCommand != nil {
		err := c.preRunCommand(c)
		if err != nil {
			return err
		}
	}

	return c.rootCommand.run(args)
}