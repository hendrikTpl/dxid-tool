package main

import (
	"log"

	"github.com/hendrikTpl/dxid-tool/pkg/util"
)

// InstallCommand handles the installation logic based on the target
func InstallCommand(target string) {
	switch target {
	case "docker":
		util.InstallDocker()
	default:
		log.Fatalf("Unsupported installation target: %s", target)
	}
}
