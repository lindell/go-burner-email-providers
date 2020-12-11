package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"
)

const emailListURL = "https://raw.githubusercontent.com/wesbos/burner-email-providers/master/emails.txt"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	// Make a request to fetch the email list
	resp, err := client.Get(emailListURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Open the list file
	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}
	file, err := os.Create(path.Join(currentPath, "burner/list.go"))
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := fmt.Fprint(file, "// Code generated (see tools/generate-list/main.go) DO NOT EDIT.\n\npackage burner\n\nvar domains = map[string]struct{}{\n"); err != nil {
		return err
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if _, err := fmt.Fprintf(file, "\t\"%s\": {},\n", scanner.Text()); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(file, "}\n"); err != nil {
		return err
	}

	return nil
}
