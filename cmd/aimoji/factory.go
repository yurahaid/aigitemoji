package aimoji

import (
	"github.com/spf13/cobra"
)

func CreateCommands() []*cobra.Command {
	return []*cobra.Command{
		NewCommitCmd(),
	}
}
