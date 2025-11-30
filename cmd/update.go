package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func updateCommand() {
	updaterURL := "https://raw.githubusercontent.com/KlangLang/loom/main/cmd/helpers/update.sh"
	l := NewLog()
	
	tmpDir, err := os.MkdirTemp("", "loom-update-*")
	if err != nil {
		fmt.Printf("%s✖%s Failed to create temp directory: %v\n", l.ERROR_COLOR, l.RESET_COLOR, err)
		return
	}
	defer os.RemoveAll(tmpDir)
	
	tmpFile := tmpDir + "/update.sh"

	curl := exec.Command("curl", "-sL", "-f", "--max-time", "10", updaterURL, "-o", tmpFile)
	if err := curl.Run(); err != nil {
		fmt.Printf("%s✖%s Failed to download update script: %v\n", l.ERROR_COLOR, l.RESET_COLOR, err)
		return
	}

	fileInfo, err := os.Stat(tmpFile)
	if err != nil || fileInfo.Size() == 0 {
		fmt.Printf("%s✖%s Downloaded file is empty or missing\n", l.ERROR_COLOR, l.RESET_COLOR)
		return
	}

	if err := os.Chmod(tmpFile, 0755); err != nil {
		fmt.Printf("%s✖%s Failed to chmod update script: %v\n", l.ERROR_COLOR, l.RESET_COLOR, err)
		return
	}

	cmd := exec.Command("bash", tmpFile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("%s✖%s Update failed: %v\n", l.ERROR_COLOR, l.RESET_COLOR, err)
		return
	}

	fmt.Printf("%s✔%s Loom updated successfully!\n", l.SUCESS_COLOR, l.RESET_COLOR)
}