package utilities

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/DylanDevelops/tmpo/internal/ui"
	"github.com/DylanDevelops/tmpo/internal/update"
	"github.com/spf13/cobra"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Long:  "Display the current version information including date and release URL.",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			DisplayVersionWithUpdateCheck()
		},
	}

	return cmd
}

func DisplayVersionWithUpdateCheck() {
	fmt.Print(GetVersionOutput())
	checkForUpdates()
}

func GetVersionOutput() string {
	versionLine := fmt.Sprintf("tmpo version %s %s", ui.Success(Version), ui.Muted(GetFormattedDate(Date)))
	changelogLine := ui.Muted(GetChangelogUrl(Version))
	return fmt.Sprintf("\n%s\n%s\n\n", versionLine, changelogLine)
}

func GetFormattedDate(inputDate string) string {
	date, err := time.Parse(time.RFC3339, inputDate)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("(%s)", date.Format("01-02-2006"))
}

func GetChangelogUrl(version string) string {
	path := "https://github.com/DylanDevelops/tmpo"

	r := regexp.MustCompile(`^v?\d+\.\d+\.\d+(-[\w.]+)?$`)
	if !r.MatchString(version) {
		return fmt.Sprintf("%s/releases/latest", path)
	}

	return fmt.Sprintf("%s/releases/tag/v%s", path, strings.TrimPrefix(version, "v"))
}

func checkForUpdates() {
	// Only check if we have a valid version (not "dev" or empty)
	if Version == "" || Version == "dev" {
		return
	}

	updateInfo, err := update.CheckForUpdate(Version)
	if err != nil {
		// Silently fail and don't bother the user with network errors
		return
	}

	if updateInfo.HasUpdate {
		fmt.Printf("%s %s\n", ui.Info("New Update Available:"), ui.Bold(strings.TrimPrefix(updateInfo.LatestVersion, "v")))
		fmt.Printf("%s\n\n", ui.Muted(updateInfo.UpdateURL))
	}
}
