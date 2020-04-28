//go:generate mockery -name CLIPrinter

package helper

import (
	"github.com/monitoror/monitoror/cli"
)

const (
	website           = "https://monitoror.com"
	developmentGuides = "https://monitoror.com/guides/#development"
	documentation     = "https://monitoror.com/" + "%s" + "documentation/"

	errorSymbol = `/!\`

	banner = `
    __  ___            _ __
   /  |/  /___  ____  (_) /_____  _________  _____
  / /|_/ / __ \/ __ \/ / __/ __ \/ ___/ __ \/ ___/
 / /  / / /_/ / / / / / /_/ /_/ / /  / /_/ / / %s
/_/  /_/\____/_/ /_/_/\__/\____/_/   \____/_/  %s

%s

`
	devModeTitle   = ` DEVELOPMENT MODE `
	uiStartCommand = `yarn serve`
	devMode        = `
┌─%s──────────────────────────────┐
│ UI must be started via %s from ./ui     │
│ For more details, check our development guide:  │
│ %s       │
└─────────────────────────────────────────────────┘

`

	monitorableHeader = `
ENABLED MONITORABLES

`

	monitorableFooterTitle = `%d more monitorables were ignored`
	monitorableFooter      = `

%s
Check the documentation to know how to enabled them:
%s
`
	echoStartup = `

Monitoror is running at:
  %s
  %s

`
)

func PrintMonitororStartupLog(monitororCli *cli.MonitororCli) {
}

func prettyPrintMonitororStartupLog() {

}

//func (p *Printer) PrintBanner() {
//	var tagFlag = ""
//	if version.BuildTags != "" {
//		tagFlag = colorer.Inverse(" " + version.BuildTags + " ")
//	}
//
//	colorer.Printf(banner, tagFlag, colorer.Green(version.Version), colorer.Blue(website))
//}
//
//func (p *Printer) PrintDevMode() {
//	colorer.Printf(devMode, colorer.Yellow(devModeTitle), colorer.Green(uiStartCommand), colorer.Blue(developmentGuides))
//}
//
//func (p *Printer) PrintMonitorableHeader() {
//	colorer.Printf(colorer.Black(colorer.Green(monitorableHeader)))
//}
//
//func (p *Printer) PrintMonitorable(displayName string, enabledVariantNames []coreModels.VariantName, erroredVariants []ErroredVariant) {
//	if len(enabledVariantNames) == 0 && len(erroredVariants) == 0 {
//		return
//	}
//
//	// Stringify variants
//	var strVariants string
//	if len(enabledVariantNames) == 1 && enabledVariantNames[0] == coreModels.DefaultVariant {
//		if len(erroredVariants) > 0 {
//			strVariants = "[default]"
//		}
//	} else {
//		var strDefault string
//		var variantsWithoutDefault []string
//
//		for _, variantName := range enabledVariantNames {
//			if variantName == coreModels.DefaultVariant {
//				strDefault = fmt.Sprintf("%s, ", variantName)
//			} else {
//				variantsWithoutDefault = append(variantsWithoutDefault, string(variantName))
//			}
//		}
//		if len(variantsWithoutDefault) > 0 {
//			strVariants = fmt.Sprintf("[%svariants: [%s]]", strDefault, strings.Join(variantsWithoutDefault, ", "))
//		}
//	}
//
//	// Print Monitorable and variants
//	prefixStatus := colorer.Green("✓")
//	if len(erroredVariants) > 0 {
//		if len(enabledVariantNames) > 0 {
//			prefixStatus = colorer.Yellow("!")
//		} else {
//			prefixStatus = colorer.Red("✕")
//		}
//	}
//	monitorableName := strings.Replace(displayName, "(faker)", colorer.Grey("(faker)"), 1)
//	colorer.Printf("  %s %s %s\n", prefixStatus, monitorableName, colorer.Grey(strVariants))
//
//	// Print Error
//	for _, erroredVariant := range erroredVariants {
//		if erroredVariant.VariantName == coreModels.DefaultVariant {
//			colorer.Printf(colorer.Red("    %s Errored %s configuration\n"), errorSymbol, erroredVariant.VariantName)
//		} else {
//			colorer.Printf(colorer.Red("    %s Errored \"%s\" variant configuration\n"), errorSymbol, erroredVariant.VariantName)
//		}
//
//		for _, err := range erroredVariant.Errors {
//			colorer.Printf("        %s\n", err.Error())
//		}
//	}
//}
//
//func (p *Printer) PrintMonitorableFooter(isProduction bool, nonEnabledMonitorableCount int) {
//	if nonEnabledMonitorableCount == 0 {
//		return
//	}
//
//	var documentationVersion string
//	if isProduction {
//		if splittedVersion := strings.Split(version.Version, "."); len(splittedVersion) == 3 {
//			documentationVersion = fmt.Sprintf("%s.%s/", splittedVersion[0], splittedVersion[1])
//		}
//	}
//
//	coloredMonitororFooterTitle := colorer.Yellow(fmt.Sprintf(monitorableFooterTitle, nonEnabledMonitorableCount))
//	colorer.Printf(
//		monitorableFooter,
//		coloredMonitororFooterTitle,
//		colorer.Blue(fmt.Sprintf(documentation, documentationVersion)),
//	)
//}
//
//func (p *Printer) PrintServerStartup(ip string, port int) {
//	colorer.Printf(
//		echoStartup,
//		colorer.Blue(fmt.Sprintf("http://localhost:%d", port)),
//		colorer.Blue(fmt.Sprintf("http://%s:%d", ip, port)),
//	)
//}
