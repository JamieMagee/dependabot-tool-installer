package install

import (
	"fmt"

	"github.com/jamiemagee/dependabot-tool-installer/internal/helpers"
	"github.com/spf13/cobra"
)

var InstallDotnetCmd = &cobra.Command{
	Use:  "dotnet",
	Long: "Install the .NET SDK",
	RunE: func(cmd *cobra.Command, args []string) error {
		d := DotnetInstaller{}

		distro, err := helpers.ReadDistro()
		if err != nil {
			return err
		}

		err = d.InstallPrerequisites(distro)
		if err != nil {
			return err
		}

		err = d.Install(args[0])
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	InstallDotnetCmd.Flags().StringP("version", "v", "latest", "Version to install")
}

type DotnetInstaller struct{}

func (d *DotnetInstaller) InstallPrerequisites(distro helpers.Distro) error {
	var err error

	switch distro.Version {
	case "24.04":
		err = helpers.AptInstall("libc6", "libgcc-s1", "libicu74", "libssl3t64", "libstdc++6", "tzdata", "zlib1g")
	}

	if err != nil {
		return err
	}

	return nil
}

func (d *DotnetInstaller) Install(version string) error {
	url := fmt.Sprintf("https://dotnetcli.azureedge.net/dotnet/Sdk/%s/dotnet-sdk-%s-linux-x64.tar.gz", version, version)
	dir, err := helpers.EnsureToolDirectory("dotnet")
	if err != nil {
		return err
	}

	err = helpers.DownloadAndExtract(url, dir)
	if err != nil {
		return err
	}

	err = helpers.LinkWrapper(dir, "dotnet", "", []string{"CLR_ICU_VERSION_OVERRIDE=74"})
	if err != nil {
		return err
	}

	return nil
}
