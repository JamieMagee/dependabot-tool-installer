package helpers

import (
	"os"
	"path"
	"strings"
)

const BaseToolDirectory = "/opt/tool"

func EnsureToolDirectory(tool string) (string, error) {
	dir := path.Join(BaseToolDirectory, tool)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func LinkWrapper(srcDir string, name string, args string, exports []string) error {
	target := path.Join("/usr/local/bin", name)

	content := "#!/usr/bin/env bash\n"

	if len(exports) > 0 {
		content += "export " + strings.Join(exports, " ") + "\n"
	}

	content += path.Join(srcDir, name) + " " + args + " \"$@\"\n"

	err := os.WriteFile(target, []byte(content), 0755)

	if err != nil {
		return err
	}

	return nil
}
