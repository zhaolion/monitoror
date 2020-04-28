//go:generate mockery -name CLIPrinter

package helper

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/monitoror/monitoror/cli"
	"github.com/monitoror/monitoror/cli/version"
	"github.com/monitoror/monitoror/internal/pkg/monitorable"
	coreModels "github.com/monitoror/monitoror/models"
	"github.com/monitoror/monitoror/pkg/system"
	"github.com/monitoror/monitoror/pkg/templates"
)

var monitororTemplate = `
    __  ___            _ __
   /  |/  /___  ____  (_) /_____  _________  _____
  / /|_/ / __ \/ __ \/ / __/ __ \/ ___/ __ \/ ___/
 / /  / / /_/ / / / / / /_/ /_/ / /  / /_/ / / {{ with .BuildTags }}{{ printf " %s " . | inverse }}{{ end }}
/_/  /_/\____/_/ /_/_/\__/\____/_/   \____/_/  {{ green .Version }}

{{ blue "https://monitoror.com" }}

{{- if .DisableUI }}

┌─ {{ yellow "DEVELOPMENT MODE" }} ──────────────────────────────┐
│ UI must be started via {{ green "yarn serve" }} from ./ui     │
│ For more details, check our development guide:  │
│ {{ blue "https://monitoror.com/guides/#development" }}       │
└─────────────────────────────────────────────────┘

{{- end }}

{{ green "ENABLED MONITORABLES" }}
{{range .MonitorableMetadata -}}
{{- if not .IsDisabled }}
{{- if not .ErroredVariantMetadata }}
  {{ green "✓" }} {{ .MonitorableName }} {{ stringifyEnabledVariantNames . | grey }}
{{- else if .EnabledVariantNames}}
  {{ yellow "!" }} {{ .MonitorableName }} {{ stringifyEnabledVariantNames . | grey }}
{{- else }}
  {{ red "x" }} {{ .MonitorableName }} {{ stringifyEnabledVariantNames . | grey }}
{{- end }}
{{- range .ErroredVariantMetadata }}
{{- if eq .VariantName "` + string(coreModels.DefaultVariant) + `" }}
    {{ printf "/!\\ Errored %s configuration" .VariantName | red }}
{{- else }}
    {{ printf "/!\\ Errored %q variant configuration" .VariantName | red }}
{{- end }}
{{- range .Errors }}
      {{ .Error }}
{{- end }}
{{- end }}
{{- end }}
{{- end }}
{{$disabledMonitorableCount := countDisabledMonitorable . -}}
{{- if ne $disabledMonitorableCount 0 }}
{{ printf "%d more monitorables were ignored" $disabledMonitorableCount | yellow }}
Check the documentation to know how to enabled them:
{{ extractDocumentationVersion .Version | printf "https://monitoror.com/%sdocumentation/" | blue }}

{{- end }}

Monitoror is running at:
  {{ printf "http://localhost:%d" .LookupPort | blue }}
  {{ printf "http://%s:%d" .LookupAddress .LookupPort | blue }}
`

type monitororInfo struct {
	Version             string
	BuildTags           string
	DisableUI           bool
	MonitorableMetadata []monitorable.Metadata
	LookupPort          int
	LookupAddress       string
}

var parsedTemplate *template.Template

func init() {
	parsedTemplate = templates.New("monitoror")
	parsedTemplate.Funcs(map[string]interface{}{
		"stringifyEnabledVariantNames": StringifyEnabledVariantNames,
		"countDisabledMonitorable":     CountDisabledMonitorable,
		"extractDocumentationVersion":  ExtractDocumentationVersion,
	})

	if _, err := parsedTemplate.Parse(monitororTemplate); err != nil {
		panic(fmt.Errorf("unable to parse monitororTemplate. %v", err))
	}
}

func PrintMonitororStartupLog(monitororCli *cli.MonitororCli) error {
	mi := &monitororInfo{
		Version:             version.Version,
		BuildTags:           version.BuildTags,
		DisableUI:           monitororCli.GetStore().CoreConfig.DisableUI,
		MonitorableMetadata: monitororCli.GetStore().MonitorableMetadata,
		LookupPort:          monitororCli.GetStore().CoreConfig.Port,
		LookupAddress:       system.GetNetworkIP(),
	}

	return parsedTemplate.Execute(monitororCli.GetOutput(), mi)
}

func StringifyEnabledVariantNames(m *monitorable.Metadata) string {
	var strVariants string
	if len(m.EnabledVariantNames) == 1 && m.EnabledVariantNames[0] == coreModels.DefaultVariant {
		if len(m.ErroredVariantMetadata) > 0 {
			strVariants = "[default]"
		}
	} else {
		var strDefault string
		var variantsWithoutDefault []string

		for _, variantName := range m.EnabledVariantNames {
			if variantName == coreModels.DefaultVariant {
				strDefault = fmt.Sprintf("%s, ", variantName)
			} else {
				variantsWithoutDefault = append(variantsWithoutDefault, string(variantName))
			}
		}
		if len(variantsWithoutDefault) > 0 {
			strVariants = fmt.Sprintf("[%svariants: [%s]]", strDefault, strings.Join(variantsWithoutDefault, ", "))
		}
	}

	return strVariants
}

func CountDisabledMonitorable(mi *monitororInfo) int {
	disabledMonitorableCount := 0
	for _, m := range mi.MonitorableMetadata {
		if m.IsDisabled() {
			disabledMonitorableCount++
		}
	}
	return disabledMonitorableCount
}

func ExtractDocumentationVersion(version string) string {
	if strings.HasSuffix(version, "-dev") {
		return ""
	}
	documentationVersion := ""
	if splittedVersion := strings.Split(version, "."); len(splittedVersion) == 3 {
		documentationVersion = fmt.Sprintf("%s.%s/", splittedVersion[0], splittedVersion[1])
	}
	return documentationVersion
}
