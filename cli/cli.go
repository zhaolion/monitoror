package cli

import (
	"github.com/spf13/cobra"

	"github.com/monitoror/monitoror/store"
)

type MonitororCli struct {
	rootCmd *cobra.Command
	store   *store.Store
}

func NewMonitororCli(store *store.Store) *MonitororCli {
	return &MonitororCli{
		store: store,
	}
}

func (c *MonitororCli) SetRootCommand(rootCmd *cobra.Command) {
	c.rootCmd = rootCmd
}

func (c *MonitororCli) GetRootCmd() *cobra.Command {
	return c.rootCmd
}

func (c *MonitororCli) GetStore() *store.Store {
	return c.store
}
