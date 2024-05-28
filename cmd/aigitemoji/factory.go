package main

import (
	"github.com/spf13/cobra"
)

func CreateCommands() []*cobra.Command {
	return []*cobra.Command{
		NewVersionCmd(),
		NewCommitCmd(),
	}
}
