package main

import (
	"context"
	"errors"
	"flag"

	"github.com/dalloriam/project/vcs"
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
	fs.StringVar(&cmd.provider, "provider", "", "Repository provider (Default defined in config)")
	fs.StringVar(&cmd.provider, "p", "", "Repository provider (Default defined in config)")
}

func (cmd *newCommand) Run(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return errors.New("must pass a project name")
	}

	cfg, err := getConfig()
	if err != nil {
		return err
	}

	var provider = cfg.PreferredProvider
	if cmd.provider != "" {
		provider = cmd.provider
	}

	svc := vcs.NewService(cfg.RootPath)
	if err := svc.Create(args[0], cfg.DefaultOwner, provider); err != nil {
		return err
	}

	return nil
}
