package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/purposed/project/vcs"
)

// TODO: Allow overriding hosts & usernames.

const (
	listCmdName = "list"
	listCmdArgs = "[-v]"
	listCmdHelp = "Lists all projects and their providers"
)

type listCommand struct {
	verbose bool
}

func (cmd *listCommand) Name() string      { return listCmdName }
func (cmd *listCommand) Args() string      { return listCmdArgs }
func (cmd *listCommand) ShortHelp() string { return listCmdHelp }
func (cmd *listCommand) LongHelp() string  { return listCmdHelp }
func (cmd *listCommand) Hidden() bool      { return false }
func (cmd *listCommand) Register(fs *flag.FlagSet) {
	fs.BoolVar(&cmd.verbose, "verbose", false, "Verbose output")
	fs.BoolVar(&cmd.verbose, "v", false, "Verbose output")
}

func (cmd *listCommand) Run(ctx context.Context, args []string) error {
	cfg, err := getConfig()
	if err != nil {
		return err
	}

	svc := vcs.NewService(cfg.RootPath)

	repos, err := svc.List(cfg.DefaultOwner)

	if err != nil {
		return err
	}

	for _, r := range repos {
		if cmd.verbose {
			fmt.Printf("* %s\n", r.Pretty())
		} else {
			fmt.Printf("* %s\n", r.Name)
		}
	}

	return nil
}
