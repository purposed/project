package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os/user"
)

const (
	cloneCmdName = "clone"
	cloneCmdHelp = "Clone a project in the appropriate directory."
	cloneCmdArgs = "REMOTE_URL"
)

type cloneCommand struct {
	rootSrcPath string
}

func (cmd *cloneCommand) Name() string      { return cloneCmdName }
func (cmd *cloneCommand) Args() string      { return cloneCmdArgs }
func (cmd *cloneCommand) ShortHelp() string { return cloneCmdHelp }
func (cmd *cloneCommand) LongHelp() string  { return cloneCmdHelp }
func (cmd *cloneCommand) Hidden() bool      { return false }
func (cmd *cloneCommand) Register(fs *flag.FlagSet) {
	fs.StringVar(&cmd.rootSrcPath, "src", "", "Root source directory (Default is '~/')")
	fs.StringVar(&cmd.rootSrcPath, "s", "", "Root source directory (Default is '~/')")
}

func (cmd *cloneCommand) Run(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return errors.New("must pass an URL to build")
	}

	repo, err := ParseRepository(args[0])
	if err != nil {
		return err
	}

	var rootPath string
	if cmd.rootSrcPath == "" {
		usr, err := user.Current()
		if err != nil {
			return err
		}
		rootPath = usr.HomeDir
	} else {
		rootPath = cmd.rootSrcPath
	}

	fmt.Printf("Cloning project [%s/%s] from %s...\n", repo.Owner, repo.Name, repo.Provider)

	fmt.Println("Ensuring all directories exist...")
	if err := repo.InitDirectories(rootPath); err != nil {
		return err
	}

	fmt.Println("Cloning...")
	if err := repo.Clone(rootPath); err != nil {
		return err
	}

	return nil
}
