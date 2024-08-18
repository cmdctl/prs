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
	switch command {
  case "--help", "-h":
    printUsage()
	case "complete":
		completePRs()
	case "abandon":
		abandonPRs()
	default:
		printUsage()
	}
}

func extractIDfromLink(link string) string {
	parts := strings.Split(link, "/")

	id := parts[len(parts)-1]
	id = strings.Split(id, "?")[0]
	return id
}

func executeComplete(id string) error {
	// command:
	// az repos pr update --id 123 --status completed --squash --delete-source-branch --merge-strategy squash
	cmd := exec.Command("az", "repos", "pr", "update", "--id", id, "--status", "completed", "--squash", "--delete-source-branch", "--merge-strategy", "squash")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("[ERROR] could not execute command complete: %s\n", out)
		return err
	}
	return nil
}

func executeAbandon(id string) error {
	// command:
	// az repos pr update --id <PR_ID> --status abandoned
	cmd := exec.Command("az", "repos", "pr", "update", "--id", id, "--status", "abandoned")
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

func completePRs() {
	onPRs(func(id string) error {
		fmt.Printf("complete PR: %s\n", id)
		err := executeAbandon(id)
		if err != nil {
			fmt.Printf("[ERROR] could not complete PR: %s error: %s", id, err)
			return err
		}
		return nil
	})
}

func abandonPRs() {
	onPRs(func(id string) error {
		fmt.Printf("abandon PR: %s\n", id)
		err := executeAbandon(id)
		if err != nil {
			fmt.Printf("[ERROR] could not abandon PR: %s error: %s", id, err)
			return err
		}
		return nil
	})
}

func printUsage() {
	fmt.Println("---------------------------")
	fmt.Println("USAGE: prs complete | abandon")
	fmt.Println("Make sure to provide a list of PRs to STDIN - each link should be on a new line")
  fmt.Println("EXAMPPLE 1: cat prs.txt | prs complete")
  fmt.Println("EXAMPPLE 2: prs complete < prs.txt")
	fmt.Println("---------------------------")
}
