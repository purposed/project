package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
)

const (
	newCmdName = "new"
	newCmdHelp = "Create a project in the appropriate repository."
	newCmdArgs = "[-p PROVIDER] PROJECT_NAME"
)

type newCommand struct {
	provider string
}

func (cmd *newCommand) Name() string      { return newCmdName }
func (cmd *newCommand) Args() string      { return newCmdArgs }
func (cmd *newCommand) ShortHelp() string { return newCmdHelp }
func (cmd *newCommand) LongHelp() string  { return newCmdHelp }
func (cmd *newCommand) Hidden() bool      { return false }
func (cmd *newCommand) Register(fs *flag.FlagSet) {
	fs.StringVar(&cmd.provider, "provider", "", "Repository provider (Default is 'github'")
	fs.StringVar(&cmd.provider, "p", "", "Repository provider (Default is 'github'")
}

func (cmd *newCommand) Run(ctx context.Context, args []string) error {
	if len(args) <= 1 {
		return errors.New("must pass a project name")
	}

	var provider = "github"
	if cmd.provider != "" {
		provider = cmd.provider
	}

	repo := Repository{
		Name:     args[0],     // TODO: Escape properly.
		Owner:    "dalloriam", // TODO: Change to env var w/ default.
		Provider: provider,
		URL:      "", // TODO: Propagate when properly created in the cloud
	}

	fmt.Println(repo)

	return nil
}
