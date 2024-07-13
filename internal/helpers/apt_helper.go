package helpers

import (
	"log/slog"
	"os/exec"
	"strings"
)

func AptInstall(pkgs ...string) error {
	todo := []string{}

	for _, pkg := range pkgs {
		installed, err := isInstalled(pkg)
		if err != nil {
			return err
		}
		if !installed {
			todo = append(todo, pkg)
		}
	}

	if len(todo) == 0 {
		return nil
	}

	cmd := exec.Command("apt-get", "-qq", "update")
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("apt-get", append([]string{"-y", "install"}, todo...)...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	slog.Debug(string(output))

	return nil
}

func isInstalled(pkg string) (bool, error) {
	cmd := exec.Command("dpkg", "-s", pkg)
	output, err := cmd.Output()
	if err != nil {
		return false, nil
	}

	return strings.Contains(string(output), "Status: install ok installed\n"), nil
}
