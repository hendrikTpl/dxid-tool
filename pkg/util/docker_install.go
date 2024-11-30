package util

import (
	"fmt"
	"log"
	"os/exec"
)

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

	for _, cmd := range commands {
		if err := executeCommand(cmd); err != nil {
			log.Fatalf("Failed to execute command: %v", err)
		}
	}

	// Step 3: Post-installation setup
	fmt.Println("Configuring Docker post-installation...")
	postInstallCommands := [][]string{
		{"sudo", "groupadd", "docker"},
		{"sudo", "usermod", "-aG", "docker", "$USER"},
		{"newgrp", "docker"},
		{"sudo", "systemctl", "enable", "docker.service"},
		{"sudo", "systemctl", "enable", "containerd.service"},
	}

	for _, cmd := range postInstallCommands {
		if err := executeCommand(cmd); err != nil {
			log.Fatalf("Failed to execute command: %v", err)
		}
	}

	fmt.Println("Verifying Docker installation...")
	if err := executeCommand([]string{"docker", "run", "hello-world"}); err != nil {
		log.Fatalf("Docker verification failed: %v", err)
	}

	fmt.Println("Docker installed successfully!")
}

func executeCommand(cmd []string) error {
	command := exec.Command(cmd[0], cmd[1:]...)
	command.Stdout = nil
	command.Stderr = nil
	return command.Run()
}
