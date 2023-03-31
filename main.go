/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"cyberland-cli/cmd"
	"embed"
)

//go:embed templates
var templateFs embed.FS

func main() {
	cmd.TemplateFs = templateFs
	cmd.Execute()
}
