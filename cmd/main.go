package main

import (
	cmd "github.com/bhatti/GSSI/cmd/commands"
)

var (
	Version = "dev"
	Commit  = "dirty"
	Date    = ""
)

func main() {
	if err := cmd.Execute(Version, Commit, Date); err != nil {
		panic(err)
	}
}
