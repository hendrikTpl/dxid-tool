package util

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// InstallDocker handles the Docker installation process
func InstallDocker() {
	fmt.Println("Starting Docker installation...")

	// Step 1: Check OS architecture and version
	arch := CheckOSArch()
	CheckUbuntuVersion()
	fmt.Printf("OS architecture: %s\n", arch)

	// Step 2: Set up Docker's apt repository
	fmt.Println("Setting up Docker's apt repository...")
	commands := [][]string{
		{"sudo", "apt-get", "update"},
		{"sudo", "apt-get", "install", "-y", "ca-certificates", "curl"},
		{"sudo", "install", "-m", "0755", "-d", "/etc/apt/keyrings"},
		{"sudo", "curl", "-fsSL", "https://download.docker.com/linux/ubuntu/gpg", "-o", "/etc/apt/keyrings/docker.asc"},
		{"sudo", "chmod", "a+r", "/etc/apt/keyrings/docker.asc"},
		{"bash", "-c", "echo \"deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo \"$VERSION_CODENAME\") stable\" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null"},
		{"sudo", "apt-get", "update"},
		{"sudo", "apt-get", "install", "-y", "docker-ce", "docker-ce-cli", "containerd.io", "docker-buildx-plugin", "docker-compose-plugin"},
	}

	for i, cmd := range commands {
		fmt.Printf("Step %d: Executing command: %v\n", i+1, cmd)
		if err := executeCommand(cmd); err != nil {
			log.Fatalf("Failed to execute command: %v\n", err)
		}
	}

	// Step 3: Post-installation setup
	fmt.Println("Configuring Docker post-installation...")
	postInstallCommands := [][]string{}

	// Check if docker group exists
	if groupExists("docker") {
		fmt.Println("Docker group already exists. Skipping group creation.")
	} else {
		fmt.Println("Creating docker group...")
		postInstallCommands = append(postInstallCommands, []string{"sudo", "groupadd", "docker"})
	}

	// Add user to the docker group (expand $USER)
	currentUser := os.Getenv("USER")
	if currentUser == "" {
		log.Fatal("Failed to detect the current user.")
	}
	fmt.Printf("Adding user '%s' to the docker group...\n", currentUser)
	postInstallCommands = append(postInstallCommands, []string{"sudo", "usermod", "-aG", "docker", currentUser})

	// Activate new group membership
	postInstallCommands = append(postInstallCommands, []string{"newgrp", "docker"})

	// Enable Docker services
	postInstallCommands = append(postInstallCommands, []string{"sudo", "systemctl", "enable", "docker.service"})
	postInstallCommands = append(postInstallCommands, []string{"sudo", "systemctl", "enable", "containerd.service"})

	for i, cmd := range postInstallCommands {
		fmt.Printf("Post-installation Step %d: Executing command: %v\n", i+1, cmd)
		if err := executeCommand(cmd); err != nil {
			log.Fatalf("Failed to execute post-installation command: %v\n", err)
		}
	}

	// Verify Docker installation
	fmt.Println("Verifying Docker installation...")
	if err := executeCommand([]string{"docker", "run", "hello-world"}); err != nil {
		log.Fatalf("Docker verification failed: %v\n", err)
	}

	fmt.Println("Docker installed successfully!")
}

// groupExists checks if a system group exists
func groupExists(group string) bool {
	cmd := exec.Command("getent", "group", group)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

// executeCommand runs a command and prints its output
func executeCommand(cmd []string) error {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = nil
	command.Stderr = nil
	return command.Run()
}
