package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/purposed/project/vcs"
)

const (
	syncCmdName = "sync"
	syncCmdArgs = "[-v]"
	syncCmdHelp = "Sync all projects by pushing pending commits & fetching all incoming changes."
)

type syncCommand struct {
	verbose bool
}

func (cmd *syncCommand) Name() string      { return syncCmdName }
func (cmd *syncCommand) Args() string      { return syncCmdArgs }
func (cmd *syncCommand) ShortHelp() string { return syncCmdHelp }
func (cmd *syncCommand) LongHelp() string  { return syncCmdHelp }
func (cmd *syncCommand) Hidden() bool      { return false }
func (cmd *syncCommand) Register(fs *flag.FlagSet) {
	fs.BoolVar(&cmd.verbose, "verbose", false, "Verbose output")
	fs.BoolVar(&cmd.verbose, "v", false, "Verbose output")
}

func (cmd *syncCommand) Run(ctx context.Context, args []string) error {
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
		fmt.Printf("* Syncing %s... ", r.Name)
		if err := r.Fetch(); err != nil {
			fmt.Printf("Error: %s.\n", err.Error())
		} else {
			fmt.Println("Done.")
		}
	}

	return nil
}
