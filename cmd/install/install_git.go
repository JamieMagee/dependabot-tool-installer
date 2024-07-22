package install

import (
	"os"
	"os/exec"

	"github.com/jamiemagee/dependabot-tool-installer/internal/helpers"
	"github.com/spf13/cobra"
)

var InstallGitCmd = &cobra.Command{
	Use:     "git",
	Long:    "Install Git",
	Example: "dependabot-tools install git",
	Args:    cobra.ExactArgs(0),
	RunE: func(_ *cobra.Command, args []string) error {
		distro, err := helpers.ReadDistro()
		if err != nil {
			return err
		}

		g := GitInstaller{}

		err = g.InstallPrerequisites(distro)
		if err != nil {
			return err
		}

		err = g.Install(distro, args)
		if err != nil {
			return err
		}

		return nil
	},
}

type GitInstaller struct {
	Installer
}

func (g GitInstaller) InstallPrerequisites(_ helpers.Distro) error {
	err := helpers.AptInstall("gnupg")
	if err != nil {
		return err
	}

	return nil
}

func (g GitInstaller) Install(_ helpers.Distro, _ []string) error {
	filePath := "/etc/apt/sources.list.d/git.list"
	contents := "deb http://ppa.launchpad.net/git-core/ppa/ubuntu noble main"
	err := os.WriteFile(filePath, []byte(contents), 0644)
	if err != nil {
		return err
	}

	cmd := exec.Command("apt-key", "adv", "--keyserver", "keyserver.ubuntu.com", "--recv-keys", "E1DD270288B4E6030699E45FA1715D88E1DF1F24")
	err = cmd.Run()

	if err != nil {
		return err
	}

	err = helpers.AptInstall("git")
	if err != nil {
		return err
	}

	return nil
}
