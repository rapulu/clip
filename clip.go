package clip



func NewCli(name, description, version string) *Cli {
	result := &Cli{
		version: version,
	}


	result.rootCommand = NewCommand(name, description)
	result.rootCommand.setApp(result)

	return result
}