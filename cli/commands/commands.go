package commands

import (
	"github.com/monitoror/monitoror/cli"
	"github.com/monitoror/monitoror/cli/commands/version"
)

func AddCommands(cli *cli.MonitororCli) {
	cli.GetRootCmd().AddCommand(
		version.NewVersionCommand(cli),
	)
}
