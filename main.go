package main

import (
	"context"
	"flag"

	"github.com/dalloriam/project/version"
	"github.com/genuinetools/pkg/cli"
)

func main() {
	p := cli.NewProgram()
	p.Description = "Software project management tool"
	p.Name = "project"
	p.Version = version.VERSION
	p.GitCommit = version.GITCOMMIT

	p.Commands = []cli.Command{
		&cloneCommand{},
		&newCommand{},
	}

	p.FlagSet = flag.NewFlagSet("project", flag.ExitOnError)
	p.Before = func(ctx context.Context) error { return nil }

	p.Run()
}
