package main

import (
	"github.com/rapulu/clip"
)



func main() {
	cli := clip.NewCli("Basic", "A basic example", "v0.01")

	cli.LongDescription("This app prints hello world")

	cli.Action(func() error {
		println("Hello world!")
		return nil
	})

	cli.Run()
}