package install

import (
	"fmt"

	"github.com/jamiemagee/dependabot-tool-installer/internal/helpers"
	"github.com/spf13/cobra"
)

var InstallDotnetCmd = &cobra.Command{
	Use:     "dotnet",
	Long:    "Install the .NET SDK",
	Example: "dependabot-tools install dotnet 8.0.100",
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

		d := DotnetInstaller{}

		err = d.InstallPrerequisites(distro)
		if err != nil {
			return err
		}

		err = d.Install(distro, arch, args)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	InstallDotnetCmd.Flags().StringP("version", "v", "latest", "Version to install")
}

type DotnetInstaller struct {
	Installer
}

// TODO: Set CLR_ICU_VERSION_OVERRIDE to the installed version
func (d DotnetInstaller) InstallPrerequisites(distro helpers.Distro) error {
	var err error

	switch distro.Name {
	case "ubuntu":
		switch distro.Version {
		case "24.04":
			err = helpers.AptInstall("libc6", "libgcc-s1", "libicu74", "libssl3t64", "libstdc++6", "tzdata", "zlib1g")
		case "22.04":
			err = helpers.AptInstall("libc6", "libgcc1", "libgssapi-krb5-2", "libicu70", "libssl3", "libstdc++6", "zlib1g")
		case "20.04":
			err = helpers.AptInstall("libc6", "libgcc1", "libgssapi-krb5-2", "libicu66", "libssl1.1", "libstdc++6", "zlib1")
		default:
			err = fmt.Errorf("unsupported Ubuntu version: %s", distro.Version)
		}
	default:
		err = fmt.Errorf("unsupported distro: %s", distro.Name)
	}

	if err != nil {
		return err
	}

	return nil
}

func (d DotnetInstaller) Install(_ helpers.Distro, _ helpers.Arch, args []string) error {
	url := fmt.Sprintf("https://dotnetcli.azureedge.net/dotnet/Sdk/%s/dotnet-sdk-%s-linux-x64.tar.gz", args[0], args[0])
	dir, err := helpers.EnsureToolDirectory("dotnet")
	if err != nil {
		return err
	}

	err = helpers.DownloadAndExtract(url, dir)
	if err != nil {
		return err
	}

	err = helpers.LinkWrapper(dir, "dotnet", "", []string{})
	if err != nil {
		return err
	}

	return nil
}
