package main

import (
	"context"
	"flag"
)

const (
	installCmdName = "install"
	installCmdHelp = "Install a binary from its official release page."
	installCmdArgs = "PROJECT_NAME_OR_PATH"
)

type installCommand struct {
}

func (cmd *installCommand) Name() string              { return installCmdName }
func (cmd *installCommand) Args() string              { return installCmdArgs }
func (cmd *installCommand) ShortHelp() string         { return installCmdHelp }
func (cmd *installCommand) LongHelp() string          { return installCmdHelp }
func (cmd *installCommand) Hidden() bool              { return false }
func (cmd *installCommand) Register(fs *flag.FlagSet) {}

func (cmd *installCommand) Run(ctx context.Context, args []string) error {
	return nil
}
