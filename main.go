package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		err := errors.New("[ERROR] you must provide a command")
		fmt.Println(err)
		printUsage()
		os.Exit(1)
	}
	command := os.Args[1]
	restOfArgs := os.Args[2:]
	switch command {
	case "--help", "-h":
		printUsage()
	case "complete":
		completePRs(restOfArgs...)
	case "abandon":
		abandonPRs(restOfArgs...)
	default:
		printUsage()
	}
}

func extractIDfromLink(link string) string {
  uri := strings.Split(link, "?")[0]
  parts := strings.Split(uri, "/")

	id := parts[len(parts)-1]
	return id
}

func executeComplete(id string, cliArgs ...string) error {
	cmdList := []string{"repos", "pr", "update", "--id", id, "--status", "completed"}

	if hasOption("--delete-source-branch", cliArgs) {
		cmdList = append(cmdList, "--delete-source-branch")
	}

	if hasOption("--squash", cliArgs) {
		cmdList = append(cmdList, "--squash")
	}

	cmd := exec.Command("az", cmdList...)
	out, err := cmd.Output()
	if err != nil {
    fmt.Printf("[ERROR] could not execute command complete: %s error: %s\n", out, err)
		return err
	}
	return nil
}

func executeAbandon(id string, cliArgs ...string) error {
	cmdList := []string{"repos", "pr", "update", "--id", id, "--status", "abandoned"}
	if hasOption("--delete-source-branch", cliArgs) {
		cmdList = append(cmdList, "--delete-source-branch")
	}
	cmd := exec.Command("az", cmdList...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("[ERROR] could not execute command abandon: %s\n", out)
		return err
	}
	return nil
}

type PRHandler func(string) error

func onPRs(handler PRHandler) {
	buf := new(bytes.Buffer)
	io.Copy(buf, os.Stdin)
	bufStr := buf.String()
	prsLinks := strings.Split(bufStr, "\n")
	for _, link := range prsLinks {
		link = strings.TrimSpace(link)
		if link == "" {
			continue
		}
		if strings.HasPrefix(link, "--") {
			continue
		}
		handler(extractIDfromLink(link))
	}
}

func hasOption(option string, options []string) bool {
	for _, opt := range options {
		if opt == option {
			return true
		}
	}
	return false
}

func completePRs(restOfArgs ...string) {
	onPRs(func(id string) error {
		fmt.Printf("complete PR: %s\n", id)
		err := executeComplete(id, restOfArgs...)
		if err != nil {
			fmt.Printf("[ERROR] could not complete PR: %s error: %s", id, err)
			return err
		}
		return nil
	})
}

func abandonPRs(restOfArgs ...string) {
	onPRs(func(id string) error {
		fmt.Printf("abandon PR: %s\n", id)
		err := executeAbandon(id, restOfArgs...)
		if err != nil {
			fmt.Printf("[ERROR] could not abandon PR: %s error: %s", id, err)
			return err
		}
		return nil
	})
}

func printUsage() {
	usage := `
Usage:
  prs [command] [options]

Commands:
  complete    Complete the specified pull requests
  abandon     Abandon the specified pull requests

Options:
  --help, -h  Show this help message and exit

Complete Command Options:
  --delete-source-branch  Delete the source branch after completing the pull request
  --squash                Squash the commits when completing the pull request

Abandon Command Options:
  --delete-source-branch  Delete the source branch after abandoning the pull request
`
	fmt.Println(usage)
}
