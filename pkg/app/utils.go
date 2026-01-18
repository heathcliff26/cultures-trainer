package app

import (
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// NOTE: The $Format strings are replaced during 'git archive' thanks to the
// companion .gitattributes file containing 'export-subst' in this same
// directory.  See also https://git-scm.com/docs/gitattributes
var gitCommit string = "$Format:%H$" // sha1 from git, output of $(git rev-parse HEAD)

func init() {
	initGitCommit()
}

func initGitCommit() {
	if strings.HasPrefix(gitCommit, "$Format") {
		var commit string
		buildinfo, _ := debug.ReadBuildInfo()
		for _, item := range buildinfo.Settings {
			if item.Key == "vcs.revision" {
				commit = item.Value
				break
			}
		}
		gitCommit = commit
	}
}

// Struct for containing the current version of the app
type Version struct {
	Name, Version, Commit, Go string
}

// Extract the version information from app
func getVersion(app fyne.App) Version {
	commit := gitCommit
	if len(commit) > 7 {
		commit = commit[:7]
	}

	metadata := app.Metadata()

	name, _ := strings.CutSuffix(metadata.Name, ".exe")
	if name == "" {
		name = filepath.Base(os.Args[0])
	}

	return Version{
		Name:    name,
		Version: "v" + metadata.Version,
		Commit:  commit,
		Go:      runtime.Version(),
	}
}

// Wrap the objects in a box with border lines
func newBorder(content ...fyne.CanvasObject) fyne.CanvasObject {
	contentContainer := container.NewThemeOverride(container.NewPadded(content...), theme.DefaultTheme())
	border := widget.NewCard("", "", contentContainer)

	return container.NewThemeOverride(border, borderTheme{})
}

// Ensures that the given string has at least the specified minimum length.
// If the string is shorter than minLength, it is padded with spaces at the end.
func minimumLengthString(str string, minLength int) string {
	if len(str) < minLength {
		str += strings.Repeat(" ", minLength-len(str))
	}
	return str
}

func hexStringToUint64(hexStr string) (uint64, error) {
	hexStr, _ = strings.CutPrefix(hexStr, "0x")
	return strconv.ParseUint(hexStr, 16, 64)
}
