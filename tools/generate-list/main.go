package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

var emailListURLs = []string{
	"https://raw.githubusercontent.com/wesbos/burner-email-providers/master/emails.txt",
	"https://raw.githubusercontent.com/disposable/disposable-email-domains/master/domains.txt",
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	domains := make(map[string]bool)

	for _, url := range emailListURLs {
		resp, err := client.Get(url)
		if err != nil {
			return fmt.Errorf("fetch %s: %w", url, err)
		}

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			domain := strings.TrimSpace(strings.ToLower(scanner.Text()))
			if domain != "" && !strings.HasPrefix(domain, "#") {
				domains[domain] = true
			}
		}
		resp.Body.Close()
	}

	sorted := make([]string, 0, len(domains))
	for d := range domains {
		sorted = append(sorted, d)
	}
	sort.Strings(sorted)

	currentPath, err := os.Getwd()
	if err != nil {
		return err
	}
	file, err := os.Create(path.Join(currentPath, "burner/list.go"))
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "// Code generated (see tools/generate-list/main.go) DO NOT EDIT.\n\npackage burner\n\nvar domains = map[string]struct{}{\n")

	for _, domain := range sorted {
		fmt.Fprintf(file, "\t\"%s\": {},\n", domain)
	}

	fmt.Fprintf(file, "}\n")

	fmt.Printf("Generated %d domains\n", len(sorted))

	return nil
}
