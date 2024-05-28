package main

import (
	"fmt"
	"runtime/debug"
	"strings"

	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints current version",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "aigitemoji version %s\n", currentVersion())
		},
	}
}

func currentVersion() string {
	if bi, ok := debug.ReadBuildInfo(); ok {
		return strings.TrimLeft(bi.Main.Version, "v")
	}

	return "dirty"
}
