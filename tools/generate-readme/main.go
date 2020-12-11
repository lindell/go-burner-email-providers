package main

import (
	"bytes"
	"html/template"
	"io/ioutil"

	increase "github.com/lindell/go-burner-email-providers/tools/calculate-increase"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

type templateData struct {
	SizeDiff string
	MemDiff  string
}

func run() error {
	size, mem, err := increase.CalculateIncrease()
	if err != nil {
		return err
	}

	data := templateData{
		SizeDiff: size,
		MemDiff:  mem,
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
