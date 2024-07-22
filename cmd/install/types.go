package install

import "github.com/jamiemagee/dependabot-tool-installer/internal/helpers"

type Installer interface {
	Install(distro helpers.Distro, args []string) error
	InstallPrerequisites(distro helpers.Distro) error
}
