/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"app01/cmd"
	"app01/setup"
	"embed"
)

// majorVersion is set via ldflags, during build
var majorVersion string

//go:embed setup/templates
var templatesFs embed.FS

func main() {
	cmd.AddConfigTemplates(setup.NewTemplates(templatesFs))
	cmd.AddVersion(majorVersion)
	cmd.Execute()
}
