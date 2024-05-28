package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yurahaid/aigitemoji/internal"
	"github.com/yurahaid/aigitemoji/internal/emojiproviders"
	"github.com/yurahaid/aigitemoji/internal/git"
	"github.com/yurahaid/aigitemoji/pkg/openai"
)

func NewCommitCmd() *cobra.Command {
	var (
		amend     bool
		openAiUrl string
	)

	var key string
	cmd := &cobra.Command{
		Short: "Create commit with a suitable emoji based on the message of the comment using AI",
		Args:  cobra.MatchAll(cobra.RangeArgs(1, 2), cobra.OnlyValidArgs),
		Use:   "commit [commit massage] [safe text to AI (optional)]",
		Example: "aigitemoji commit 'Fix critical bug' \n" +
			"aigitemoji commit 'Fix critical bug 33121 in our security system' 'Fix critical bug' \n",
		Run: func(cmd *cobra.Command, args []string) {
			g, err := git.SetupGit()
			if err != nil {
				cmd.PrintErrln(err)
			}
			openAiKey := viper.GetString("open-ai-api-key")
			httpClient := &http.Client{}
			openAiClient := openai.NewClient(
				httpClient,
				openAiUrl,
				openAiKey,
				openai.Model35turbo,
			)
			emojiProvider := emojiproviders.NewChatGpt(openAiClient)
			aiGitEmoji := internal.NewAIGitEmoji(emojiProvider, g)
			commit := args[0]
			aiPrompt := args[0]
			if len(args) > 1 {
				aiPrompt = args[1]
			}

			emojiCommit, hash, err := aiGitEmoji.Commit(cmd.Context(), commit, aiPrompt, amend)
			if err != nil {
				cmd.PrintErrln(err)
			}

			cmd.Println(emojiCommit)
			cmd.Println(hash)

		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&amend, "amend", false, "Replace the tip of the current branch by creating a new commit.")
	flags.StringVar(&openAiUrl, "open-ai-url", "https://api.openai.com", "open-ai url")
	flags.StringVar(&key, "open-ai-api-key", "", "username, facultative if you have config file")
	if err := viper.BindPFlag("open-ai-api-key", flags.Lookup("open-ai-api-key")); err != nil {
		fmt.Println(err)
	}

	return cmd
}
