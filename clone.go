package main

import (
	"context"
	"errors"
	"flag"

	"github.com/dalloriam/project/vcs"
)

const (
	cloneCmdName = "clone"
	cloneCmdHelp = "Clone a project in the appropriate directory."
	cloneCmdArgs = "REMOTE_URL"
)

type cloneCommand struct {
}

func (cmd *cloneCommand) Name() string              { return cloneCmdName }
func (cmd *cloneCommand) Args() string              { return cloneCmdArgs }
func (cmd *cloneCommand) ShortHelp() string         { return cloneCmdHelp }
func (cmd *cloneCommand) LongHelp() string          { return cloneCmdHelp }
func (cmd *cloneCommand) Hidden() bool              { return false }
func (cmd *cloneCommand) Register(fs *flag.FlagSet) {}

func (cmd *cloneCommand) Run(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return errors.New("must pass an URL to build")
	}

	cfg := getConfig()

	svc := vcs.NewService(cfg.RootPath)
	if err := svc.Clone(args[0]); err != nil {
		return err
	}

	return nil
}
