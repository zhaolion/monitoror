package version

import (
	"os"
	"runtime"
	"text/tabwriter"
	"text/template"

	"github.com/monitoror/monitoror/cli"
	"github.com/monitoror/monitoror/cli/version"
	"github.com/monitoror/monitoror/pkg/templates"

	"github.com/spf13/cobra"
)

var versionTemplate = ` Version:    {{green .Version}}{{ if .BuildTags }}{{grey (printf " (%s)" .BuildTags)}}{{end}}
 Git commit: {{green .GitCommit}}
 Built:      {{green .BuildTime}}

 Go version: {{blue .GoVersion}}
 OS/Arch:    {{blue .Os}}/{{blue .Arch}}`

type versionInfo struct {
	Version   string
	GitCommit string
	GoVersion string
	Os        string
	Arch      string
	BuildTime string
	BuildTags string
}

func NewVersionCommand(monitororCli *cli.MonitororCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show the Monitoror version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runVersion(monitororCli)
		},
	}
	return cmd
}

func runVersion(_ *cli.MonitororCli) error {
	vi := versionInfo{
		Version:   version.Version,
		GitCommit: version.GitCommit,
		BuildTime: version.BuildTime,
		BuildTags: version.BuildTags,
		GoVersion: runtime.Version(),
		Os:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}

	tmpl, err := templates.New("version").Parse(versionTemplate)
	if err != nil {
		return err
	}

	return prettyPrintVersion(vi, tmpl)
}

func prettyPrintVersion(vi versionInfo, tmpl *template.Template) error {
	t := tabwriter.NewWriter(os.Stdout, 20, 1, 1, ' ', 0)
	err := tmpl.Execute(t, vi)
	_, _ = t.Write([]byte("\n"))
	_ = t.Flush()
	return err
}
