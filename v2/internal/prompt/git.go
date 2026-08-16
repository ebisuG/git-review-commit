package prompt

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func readGitLog() (string, error) {
	cmd := exec.Command("git", "log", "--pretty=format:subject:%s%nbody:%b", "-10")

	out, err := cmd.Output()
	if err != nil {
		fmt.Println("could not run command: ", err)
	}
	return string(out), nil
}

type Diff struct {
	Diff         string
	ChangedFiles []ChangedFile
}

type ChangedFile struct {
	Path     string
	AllLines string
}

func readChangesInStagedFile() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func findChangedFiles() ([]string, error) {
	cmd := exec.Command("git", "diff", "--staged", "--name-only")

	out, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}

	files := strings.Split(strings.ReplaceAll(string(out), "\r\n", "\n"), "\n")
	files = files[:len(files)-1] //Trim last empty line

	return files, nil
}

func getGitRootPath() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	rootPath := strings.Split(strings.ReplaceAll(string(out), "\r\n", "\n"), "\n")

	return rootPath[0], nil

}

func buildDiff() (Diff, error) {
	diff, err := readChangesInStagedFile()
	if err != nil {
		return Diff{}, err
	}

	changedFilesPath, err := findChangedFiles()
	if err != nil {
		return Diff{}, err
	}

	gitRoot, err := getGitRootPath()
	if err != nil {
		fmt.Println("Not Found Git Root Directory")
	}

	var changedFiles []ChangedFile
	for _, v := range changedFilesPath {
		path := filepath.Join(gitRoot, v)
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("Not Found : ", path)
		}
		changedFiles = append(changedFiles, ChangedFile{Path: v, AllLines: string(content)})
	}
	return Diff{Diff: diff, ChangedFiles: changedFiles}, nil
}
