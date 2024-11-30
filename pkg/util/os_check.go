package util

import (
	"fmt"
	"os/exec"
	"strings"
)

// CheckOSArch verifies the architecture and OS compatibility
func CheckOSArch() string {
	cmd := exec.Command("uname", "-m")
	output, err := cmd.Output()
	if err != nil {
		panic(fmt.Sprintf("Failed to detect OS architecture: %v", err))
	}

	arch := strings.TrimSpace(string(output))
	if arch != "x86_64" && arch != "amd64" && arch != "arm64" {
		panic(fmt.Sprintf("Unsupported architecture: %s", arch))
	}
	return arch
}

// CheckUbuntuVersion checks the Ubuntu version compatibility
func CheckUbuntuVersion() {
	cmd := exec.Command("lsb_release", "-cs")
	output, err := cmd.Output()
	if err != nil {
		panic(fmt.Sprintf("Failed to detect Ubuntu version: %v", err))
	}

	version := strings.TrimSpace(string(output))
	supportedVersions := []string{"oracular", "noble", "jammy", "focal"}

	for _, v := range supportedVersions {
		if version == v {
			return
		}
	}
	panic(fmt.Sprintf("Unsupported Ubuntu version: %s", version))
}
