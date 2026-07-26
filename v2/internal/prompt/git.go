package prompt

import (
	"fmt"
	"os/exec"
)

// $ git log --pretty=format:"%s%n%b-----"
// Merge pull request #13 from ebisuG/12-refactor-to-put-review-feature-on-the-center-of-the-tool
// Refactor: Separate features into small modules and improve application composition in main.go-----
// chore:rename NewPrompt to BuildPrompt
// -----
// refactor: separate features into small modules
// Design main.go as a pipeline that calls loader, reviewer, and prompt in that order. main.go doesn't need to know about the CLI; it should focus on input for building the prompt.
// -----
// Merge pull request #10 from ebisuG/9-load-the-api-key-from-a-local-configuration-file
// Load the api key from a local configuration file-----
// refactor: create new variable with explicit assignment
// Follow the preferred style in Go.
// -----

// $ git log --pretty=format:"subject:%s%nbody:%b"
// subject:Merge pull request #13 from ebisuG/12-refactor-to-put-review-feature-on-the-center-of-the-tool
// body:Refactor: Separate features into small modules and improve application composition in main.go
// subject:chore:rename NewPrompt to BuildPrompt
// body:
// subject:refactor: separate features into small modules
// body:Design main.go as a pipeline that calls loader, reviewer, and prompt in that order. main.go doesn't need to know about the CLI; it should focus on input for building the prompt.

// subject:Merge pull request #10 from ebisuG/9-load-the-api-key-from-a-local-configuration-file
// body:Load the api key from a local configuration file
// subject:refactor: create new variable with explicit assignment
// body:Follow the preferred style in Go.

// run shell
func readGitLog() (string, error) {
	cmd := exec.Command("git", "log", "--pretty=format:'subject:%s%nbody:%b'", "-10")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	// fmt.Println("Output: ", string(out))
	return string(out), nil
}
