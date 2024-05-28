package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	rootCmd := &cobra.Command{
		Version: currentVersion(),
		Use:     "aigitemoji",
		Short:   "Create commit with a suitable emojiproviders based on the message of the comment using AI",
	}

	rootCmd.AddCommand(CreateCommands()...)

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
