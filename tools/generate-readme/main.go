package main

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"regexp"

	increase "github.com/lindell/go-burner-email-providers/tools/calculate-increase"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type templateData struct {
	SizeDiff  string
	MemDiff   string
	NoDomains string
}

// Number of extra lines in list.go
const extraListLines = 8

func run() error {
	size, mem, err := increase.CalculateIncrease()
	if err != nil {
		return err
	}

	// Count lines in list.go
	file, _ := os.Open("./burner/list.go")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}

	data := templateData{
		SizeDiff:  size,
		MemDiff:   mem,
		NoDomains: formatCommas(lineCount - extraListLines),
	}

	tmpl, err := template.ParseFiles("./docs/README.template.md")
	if err != nil {
		return err
	}

	tmplBuf := &bytes.Buffer{}
	err = tmpl.Execute(tmplBuf, data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("README.md", tmplBuf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

func formatCommas(num int) string {
	str := fmt.Sprintf("%d", num)
	re := regexp.MustCompile(`(\d+)(\d{3})`)
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1,$2")
	}
	return str
}
