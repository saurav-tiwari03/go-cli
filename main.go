package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

)

func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func simpleAdd(scanner *bufio.Scanner) {
	fmt.Print("Enter commit message: ")
	scanner.Scan()
	commitMessage := scanner.Text()

	fmt.Print("Enter branch name to push (e.g., main): ")
	scanner.Scan()
	branchName := scanner.Text()

	// Run git add command and show output
	output, err := runCommand("git", "add", ".")
	if err != nil {
		fmt.Println("Error during git add:", err)
		fmt.Println(output) // Show command output (stderr)
		return
	}
	fmt.Println("Added!", output) // Show output after successful execution

	// Run git commit command and show output
	output, err = runCommand("git", "commit", "-m", commitMessage)
	if err != nil {
		fmt.Println("Error during git commit:", err)
		fmt.Println(output) // Show command output (stderr)
		return
	}
	fmt.Println("Committed!", output) // Show output after successful execution

	// Run git push command and show output
	output, err = runCommand("git", "push", "origin", branchName)
	if err != nil {
		fmt.Println("Error while pushing the code:", err)
		fmt.Println(output) // Show command output (stderr)
		return
	}
	fmt.Println("Pushed successfully!", output) // Show output after successful execution
}


func addSSHAgent(scanner *bufio.Scanner) {
	// Initialize SSH agent and capture environment variables
	output, err := runCommand("ssh-agent", "-s")
	if err != nil {
		fmt.Println("Error initializing ssh-agent:", err)
		fmt.Println(output) // Print the output even if there's an error
		return
	}
	// Show the output after running the command
	fmt.Println("SSH Agent initialized:", output)

	// Capture environment variables for SSH agent
	envVars := strings.Split(output, "\n")
	for _, envVar := range envVars {
		if strings.HasPrefix(envVar, "SSH_AUTH_SOCK") || strings.HasPrefix(envVar, "SSH_AGENT_PID") {
			parts := strings.SplitN(envVar, "=", 2)
			if len(parts) == 2 {
				// Set the environment variables for the Go process
				os.Setenv(parts[0], parts[1])
			}
		}
	}

	// Ask for the SSH key filename and add it to the agent
	fmt.Println("Please enter your ssh-agent filename you want to add!")
	scanner.Scan() // Read the input
	filename := strings.TrimSpace(scanner.Text())

	// Get the home directory and build the full path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}
	sshKeyPath := fmt.Sprintf("%s/.ssh/%s", homeDir, filename)

	// Add the SSH key using ssh-add
	output, err = runCommand("ssh-add", sshKeyPath)
	if err != nil {
		fmt.Println("Error adding SSH key:", err)
		fmt.Println(output) // Show the output in case of error
		return
	}
	fmt.Println("SSH Key added successfully:", output) // Show the output after successful command execution
}

func operation(request string, scanner *bufio.Scanner) {
	switch request {
	case "1":
		simpleAdd(scanner)
	case "2":
		addSSHAgent(scanner)
		simpleAdd(scanner)
	default:
		fmt.Println("Invalid option selected.")
	}
}

func selectOperation(scanner *bufio.Scanner) {
	fmt.Println("Choose an operation:")
	fmt.Println("1. Push without SSH")
	fmt.Println("2. Push with SSH (Initiate SSH-Agent)")
	fmt.Print("Enter choice: ")
	scanner.Scan()
	request := strings.TrimSpace(scanner.Text())
	operation(request, scanner)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	selectOperation(scanner)
}
