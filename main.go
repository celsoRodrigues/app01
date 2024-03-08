/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"app01/cmd"
)

var majorVersion string

func main() {
	cmd.AddVersion(majorVersion)
	cmd.Execute()
}
