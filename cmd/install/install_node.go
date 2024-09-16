package install

import (
	"fmt"

	"github.com/jamiemagee/dependabot-tool-installer/internal/helpers"
	"github.com/spf13/cobra"
)

var InstallNodeCmd = &cobra.Command{
	Use:     "node",
	Long:    "Install Node.js",
	Example: "dependabot-tools install node",
	Args:    cobra.ExactArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		distro, err := helpers.ReadDistro()
		if err != nil {
			return err
		}

		arch, err := helpers.FindArch()
		if err != nil {
			return err
		}

		n := NodeInstaller{}

		err = n.Install(distro, arch, args)
		if err != nil {
			return err
		}

		return nil
	},
}

type NodeInstaller struct {
	Installer
}

func (n NodeInstaller) Install(distro helpers.Distro, arch helpers.Arch, args []string) error {
	nodeArch, err := NodeArch(arch)
	if err != nil {
		return err
	}

	file := fmt.Sprintf("node-v%s-linux-%s", args[0], nodeArch)
	url := fmt.Sprintf("https://nodejs.org/dist/v%s/%s.tar.gz", args[0], file)
	dir, err := helpers.EnsureToolDirectory("node")
	if err != nil {
		return err
	}

	err = helpers.DownloadAndExtract(url, dir)
	if err != nil {
		return err
	}

	bin := fmt.Sprintf("%s/%s/bin", dir, file)
	// TODO link npm
	err = helpers.LinkWrapper(bin, "node", "", []string{})
	if err != nil {
		return err
	}

	return nil
}

func NodeArch(arch helpers.Arch) (string, error) {
	switch arch {
	case 1:
		return "x64", nil
	case 2:
		return "arm64", nil
	default:
		return "", fmt.Errorf("unsupported architecture: %s", string(arch))
	}
}
