package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
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
	output, err := runCommand("git", "add", ".")
	if err != nil {
		fmt.Println("Error during git add:", err)
		fmt.Println(output)
		return
	}
	fmt.Println("Added!", output)
	output, err = runCommand("git", "commit", "-m", commitMessage)
	if err != nil {
		fmt.Println("Error during git commit:", err)
		fmt.Println(output) 
		return
	}
	fmt.Println("Committed!", output) 
	output, err = runCommand("git", "push", "origin", branchName)
	if err != nil {
		fmt.Println("Error while pushing the code:", err)
		fmt.Println(output) 
		return
	}
	fmt.Println("Pushed successfully!", output)
}
func addSSHAgent(scanner *bufio.Scanner) {
	output, err := runCommand("ssh-agent", "-s")
	if err != nil {
		fmt.Println("Error initializing ssh-agent:", err)
		fmt.Println(output)
		return
	}

	fmt.Println("SSH Agent initialized:", output)

	re := regexp.MustCompile(`^(\w+)=([^;]+)`)
	envVars := strings.Split(output, "\n")
	for _, envVar := range envVars {
		envVar = strings.TrimSpace(envVar)
		matches := re.FindStringSubmatch(envVar)
		if matches != nil {
			key, value := matches[1], matches[2]
			os.Setenv(key, value)
			fmt.Printf("Set %s=%s\n", key, value) // Debug log
		}
	}

	// Rest of the function remains the same
	fmt.Println("Please enter your ssh-agent filename you want to add!")
	scanner.Scan()
	filename := strings.TrimSpace(scanner.Text())

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}
	sshKeyPath := fmt.Sprintf("%s/.ssh/%s", homeDir, filename)

	output, err = runCommand("ssh-add", sshKeyPath)
	if err != nil {
		fmt.Println("Error adding SSH key:", err)
		fmt.Println(output)
		return
	}
	fmt.Println("SSH Key added successfully:", output)
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
	fmt.Println("2. Push with SSH (Ensure ssh-agent is running in your terminal)")
	fmt.Print("Enter choice: ")
	scanner.Scan()
	request := strings.TrimSpace(scanner.Text())
	operation(request, scanner)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	selectOperation(scanner)
}
