package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jamiemagee/dependabot-tool-installer/cmd/install"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i"},
}

func init() {
	installCmd.AddCommand(install.InstallDotnetCmd)
	installCmd.AddCommand(install.InstallGitCmd)
	installCmd.AddCommand(install.InstallNodeCmd)
}
