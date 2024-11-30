package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define CLI arguments
	installCmd := flag.NewFlagSet("install", flag.ExitOnError)
	installTarget := installCmd.String("target", "", "Specify the target to install (e.g., 'docker')")

	if len(os.Args) < 2 {
		fmt.Println("Usage: dxid <command> [arguments]")
		fmt.Println("Commands: install")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "install":
		installCmd.Parse(os.Args[2:])
		if *installTarget == "" {
			log.Fatal("Please specify a target using --target")
		}
		InstallCommand(*installTarget) // Call the InstallCommand function from install.go
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
