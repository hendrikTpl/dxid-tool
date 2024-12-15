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

	// Define validate subcommand
	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)
	validateRTSPURL := validateCmd.String("url", "", "RTSP URL to validate (e.g., rtsp://user:pass@ip:port/path)")

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
	case "validate":
		validateCmd.Parse(os.Args[2:])
		if *validateRTSPURL == "" {
			log.Fatal("Please specify an RTSP URL using --url")
		}
		ValidateCommand(*validateRTSPURL)

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
	fmt.Println("  validate      Validate an RTSP stream")

	fmt.Println("\nGlobal Flags:")
	fmt.Println("  --version     Show the version of the tool")
	fmt.Println("  --help        Show this help message")

	fmt.Println("\nExamples:")
	fmt.Println("  dxid --version")
	fmt.Println("  dxid install --target docker")
	fmt.Println("  dxid validate --url rtsp://user:pass@ip:port/path")
}

// ValidateCommand handles the validate subcommand
func ValidateCommand(rtspURL string) {
	fmt.Printf("Validating RTSP stream: %s...\n", rtspURL)
	// Call validation logic from util package
	result, err := ValidateRTSP(rtspURL)
	if err != nil {
		log.Fatalf("RTSP validation failed: %v\n", err)
	}
	fmt.Println("RTSP validation successful!")
	fmt.Printf("Resolution: %s\n", result.Resolution)
	fmt.Printf("Codec: %s\n", result.Codec)
	fmt.Printf("Frame captured at: %s\n", result.FramePath)
}
