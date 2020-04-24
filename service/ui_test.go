package service

import (
	"testing"

	cliMocks "github.com/monitoror/monitoror/cli/helper/mocks"

	"github.com/monitoror/monitoror/config"
	"github.com/monitoror/monitoror/store"
)

func TestInitUI_Dev(t *testing.T) {
	cliMock := new(cliMocks.CLIPrinter)
	cliMock.On("PrintDevMode")
	InitUI(&Server{
		Echo: nil,
		store: &store.Store{
			CoreConfig: &config.Config{Env: "develop"},
			CliHelper:  cliMock,
		},
	})
	cliMock.AssertNumberOfCalls(t, "PrintDevMode", 1)
}
