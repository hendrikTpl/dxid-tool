package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const version = "1.0.0" // Define the global version of the CLI tool

func main() {
	// Global flags for version and help
	versionFlag := flag.Bool("version", false, "Show the version of the tool")
	helpFlag := flag.Bool("help", false, "Show help and usage information")

	// Define install subcommand
	installCmd := flag.NewFlagSet("install", flag.ExitOnError)
	installTarget := installCmd.String("target", "", "Specify the target to install (e.g., 'docker')")

	// Parse global flags
	flag.Parse()

	// Handle global flags
	if *versionFlag {
		fmt.Println("dxid version", version)
		os.Exit(0)
	}

	if *helpFlag || len(os.Args) < 2 {
		showHelp()
		os.Exit(0)
	}

	// Handle subcommands
	switch os.Args[1] {
	case "install":
		installCmd.Parse(os.Args[2:])
		if *installTarget == "" {
			log.Fatal("Please specify a target using --target")
		}
		InstallCommand(*installTarget)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		showHelp()
		os.Exit(1)
	}
}

// showHelp displays help information for the CLI tool
func showHelp() {
	fmt.Println("Usage: dxid <command> [options]")
	fmt.Println("\nCommands:")
	fmt.Println("  install       Install software (e.g., Docker)")

	fmt.Println("\nGlobal Flags:")
	fmt.Println("  --version     Show the version of the tool")
	fmt.Println("  --help        Show this help message")

	fmt.Println("\nExamples:")
	fmt.Println("  dxid --version")
	fmt.Println("  dxid install --target docker")
}
