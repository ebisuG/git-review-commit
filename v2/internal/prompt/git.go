package prompt

import (
	"fmt"
	"os/exec"
)

func readGitLog() (string, error) {
	cmd := exec.Command("git", "log", "--pretty=format:subject:%s%nbody:%b", "-10")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	return string(out), nil
}
