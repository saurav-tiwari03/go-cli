package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

)
func operation(request string) {
	if request == "1" {
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Enter commit message: ")
		scanner.Scan()
		commitMessage := scanner.Text()

		fmt.Print("Enter branch name to push (e.g., main): ")
		scanner.Scan()
		branchName := scanner.Text()

		exec.Command("git", "add", ".").CombinedOutput()
		fmt.Println("Added!")

		exec.Command("git", "commit", "-m", commitMessage).CombinedOutput()
		fmt.Println("Commited!")

		exec.Command("git", "push", "origin", branchName).CombinedOutput()
		fmt.Println("Pushed!")
	}
}

func selectOperation() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Choose an operation:")
	fmt.Println("1. Push without SSH")
	fmt.Println("2. Push with SSH (First time)")
	fmt.Println("3. Push with SSH (Subsequent times)")
	scanner.Scan()
	request := scanner.Text()

	operation(request)
}

func main() {
	selectOperation()
}
