package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

)

func operation(request string) {
	scanner := bufio.NewScanner(os.Stdin)
	switch request {
	case "1":
		fmt.Print("Enter commit message: ")
		scanner.Scan()
		commitMessage := scanner.Text()

		fmt.Print("Enter branch name to push (e.g., main): ")
		scanner.Scan()
		branchName := scanner.Text()

		if _, err := exec.Command("git", "add", ".").CombinedOutput(); err != nil {
			fmt.Println("Error during git add:", err)
			return
		}
		fmt.Println("Added!")

		if _, err := exec.Command("git", "commit", "-m", commitMessage).CombinedOutput(); err != nil {
			fmt.Println("Error during git commit:", err)
			return
		}
		fmt.Println("Commited!")

		output, err := exec.Command("git", "push", "origin", branchName).CombinedOutput()
		if err != nil {
			fmt.Println("Error while pushing the code:", err)
			fmt.Println(string(output))
			return
		}
		fmt.Println("Pushed successfully!")

	case "2":
		fmt.Print("Enter SSH repository URL (e.g., git@github.com:user/repo.git): ")
		scanner.Scan()
		repoURL := scanner.Text()

		fmt.Print("Enter commit message: ")
		scanner.Scan()
		commitMessage := scanner.Text()
		fmt.Print("Enter branch name to push (e.g., main): ")
		scanner.Scan()
		branchName := scanner.Text()
		exec.Command("git", "remote", "remove", "origin").CombinedOutput()
		if _, err := exec.Command("git", "remote", "add", "origin", repoURL).CombinedOutput(); err != nil {
			fmt.Println("Error setting remote URL:", err)
			return
		}
		fmt.Println("Remote SSH URL added successfully!")

		if _, err := exec.Command("git", "add", ".").CombinedOutput(); err != nil {
			fmt.Println("Error during git add:", err)
			return
		}
		fmt.Println("Files added successfully!")

		if _, err := exec.Command("git", "commit", "-m", commitMessage).CombinedOutput(); err != nil {
			fmt.Println("Error during git commit:", err)
			return
		}
		fmt.Println("Commit created successfully!")

		output, err := exec.Command("git", "push", "-u", "origin", branchName).CombinedOutput()
		if err != nil {
			fmt.Println("Error while pushing via SSH:", err)
			fmt.Println(string(output))
			return
		}
		fmt.Println("Pushed successfully via SSH!")

	case "3":
		fmt.Print("Enter commit message: ")
		scanner.Scan()
		commitMessage := scanner.Text()

		fmt.Print("Enter branch name to push (e.g., main): ")
		scanner.Scan()
		branchName := scanner.Text()

		if _, err := exec.Command("git", "add", ".").CombinedOutput(); err != nil {
			fmt.Println("Error during git add:", err)
			return
		}
		fmt.Println("Added!")

		if _, err := exec.Command("git", "commit", "-m", commitMessage).CombinedOutput(); err != nil {
			fmt.Println("Error during git commit:", err)
			return
		}
		fmt.Println("Commit created successfully!")

		output, err := exec.Command("git", "push", "origin", branchName).CombinedOutput()
		if err != nil {
			fmt.Println("Error while pushing via SSH:", err)
			fmt.Println(string(output))
			return
		}
		fmt.Println("Pushed successfully via SSH!")

	default:
		fmt.Println("Invalid option selected.")
	}
}

func selectOperation() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose an operation:")
	fmt.Println("1. Push without SSH")
	fmt.Println("2. Push with SSH (First time)")
	fmt.Println("3. Push with SSH (Subsequent times)")
	fmt.Print("Enter choice: ")
	scanner.Scan()
	request := strings.TrimSpace(scanner.Text())

	operation(request)
}

func main() {
	selectOperation()
}
