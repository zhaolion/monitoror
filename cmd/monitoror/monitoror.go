package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/monitoror/monitoror/cli"
	"github.com/monitoror/monitoror/cli/commands"
	"github.com/monitoror/monitoror/cli/helper"
	"github.com/monitoror/monitoror/cli/version"
	"github.com/monitoror/monitoror/config"
	"github.com/monitoror/monitoror/registry"
	"github.com/monitoror/monitoror/service"
	"github.com/monitoror/monitoror/store"

	"github.com/joho/godotenv"
	"github.com/jsdidierlaurent/echo-middleware/cache"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newMonitororCommand(monitororCli *cli.MonitororCli) {
	cmd := &cobra.Command{
		Use:   "monitoror",
		Short: "Unified monitoring wallboard",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Init Service
			server := service.Init(monitororCli.GetStore())

			if err := helper.PrintMonitororStartupLog(monitororCli); err != nil {
				return err
			}
			return server.Start()
		},
		Version: fmt.Sprintf("%s, build %s", version.Version, version.GitCommit),
	}

	cmd.Flags().BoolP("version", "v", false, "Print version information and quit")
	cmd.PersistentFlags().BoolP("debug", "d", false, "Start monitoror in debug mode")
	_ = viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))

	// Setup this command as root command in cli
	monitororCli.SetRootCommand(cmd)

	// add other command
	commands.AddCommands(monitororCli)
}

func main() {
	// Setup logger
	log.SetPrefix("")
	log.SetHeader("[${level}]")
	log.SetLevel(log.INFO)

	// Load .env
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	_ = godotenv.Load(".env")
	_ = godotenv.Load(filepath.Join(dir, ".env"))

	// Setup Store
	store := &store.Store{
		CoreConfig: config.InitConfig(),
		Registry:   registry.NewRegistry(),
		CacheStore: cache.NewGoCacheStore(time.Minute*5, time.Second), // Default value, Always override
	}

	// Init CLI
	monitororCli := cli.NewMonitororCli(store)
	newMonitororCommand(monitororCli)

	// Start CLI
	if err := monitororCli.GetRootCommand().Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
