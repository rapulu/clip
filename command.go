package clip

import (
	"flag"
	"fmt"
	"os"
)


type Command struct {
	name 							string
	shortdescription 	string
	longdescription 	string
	actionCallback 		Action
	subCommands 			[]*Command
	subCommandsMap 		map[string]*Command
	flags 						*flag.FlagSet
	flagCount 				int
	app 							*Cli
	commandPath				string
}


func NewCommand(name, description string) *Command {
	result := &Command{
		name:             name,
		shortdescription: description,
	}

	return result
}

func (c *Command) Action(callback Action) *Command {
	c.actionCallback = callback
	return c
}

func (c *Command) parseFlags(args []string) error {
	tmp := os.Stderr
	os.Stderr = nil
	err := c.flags.Parse(args)
	os.Stderr = tmp
	return err
}

func (c *Command) run(args []string) error {

	//if we have argumnets process them
	if len(args) > 0 {
		//Check for subcommand
		subcommand := c.subCommandsMap[args[0]]
		if subcommand != nil {
			return subcommand.run(args[1:])
		}

		//Parse flags
		err := c.parseFlags(args)
		if err != nil {
			if c.app.errorHandler != nil {
				return c.app.errorHandler(c.commandPath, err)
			}
			return fmt.Errorf("Error: %s\nSee '%s --help' for usage", err, c.commandPath)
		}

	}

	if c.actionCallback != nil {
		c.actionCallback()
	}

	return nil
}

func (c *Command) setApp(app *Cli) {
	c.app = app
}