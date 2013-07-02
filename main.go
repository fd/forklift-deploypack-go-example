package main

import (
	"bitbucket.org/mrhenry/forklift/deploypack/command"
)

func main() { command.Run(&handler{}) }

type handler struct{}

func (h *handler) ProcessConfiguration(ctx *command.Context) error {
	var (
		config interface{}
		err    error
	)

	err = ctx.LoadConfig(&config)
	if err != nil {
		return err
	}

	// Do something with the config here.

	err = ctx.DumpConfig(config)
	if err != nil {
		return err
	}

	return nil
}
