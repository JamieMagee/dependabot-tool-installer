package helpers

import (
	"fmt"
	"os/exec"
	"strings"
)

type Arch int

const (
	_ Arch = iota
	x64
	arm64
)

func FindArch() (Arch, error) {
	cmd := exec.Command("uname", "-m")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	switch strings.TrimSpace(string(output)) {
	case "x86_64":
		return x64, nil
	case "aarch64":
		return arm64, nil
	default:
		return 0, fmt.Errorf("unsupported architecture: %s", string(output))
	}

}
