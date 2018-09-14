package main

import (
	"bufio"
	"fmt"
	"github.com/willis7/gist-diff/github"
	"os"
)

func main() {
	client := github.New("config.json")


	// take input from Stdin
	input := bufio.NewScanner(os.Stdin)
	var lines []string
	for input.Scan() {
		lines = append(lines, input.Text())
	}

	// TODO: remove debug
	for _, line := range lines {
		fmt.Println(line)
	}

	// remove garbage

	// connect to github

	// push to github
	client.Publish(lines)
}
