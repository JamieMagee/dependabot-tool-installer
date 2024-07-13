package install

import "github.com/spf13/cobra"

var InstallGitCmd = &cobra.Command{
	Use:  "git",
	Args: cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		d := GitInstaller{}

		err := d.InstallPrerequisites()
		if err != nil {
			return err
		}

		err = d.Install()
		if err != nil {
			return err
		}

		return nil
	},
}

type GitInstaller struct{}

func (d *GitInstaller) InstallPrerequisites() error {
	return nil
}

func (d *GitInstaller) Install() error {
	return nil
}
