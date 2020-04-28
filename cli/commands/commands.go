package commands

import (
	"github.com/monitoror/monitoror/cli"
	"github.com/monitoror/monitoror/cli/commands/version"
)

func AddCommands(cli *cli.MonitororCli) {
	cli.GetRootCommand().AddCommand(
		version.NewVersionCommand(cli),
	)
}
