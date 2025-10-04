package main

import (
	"embed"

	"github.com/kangjoa/gostart/cmd"
)

//go:embed templates/*
var templateFiles embed.FS

func main() {
	cmd.Execute(templateFiles)
}
