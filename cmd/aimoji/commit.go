package aimoji

import (
	"fmt"
	"net/http"

	"aigitemoji/internal"
	"aigitemoji/internal/emojiproviders"
	"aigitemoji/internal/git"
	"aigitemoji/pkg/openai"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCommitCmd() *cobra.Command {
	var (
		amend     bool
		openAiUrl string
	)

	var token string
	cmd := &cobra.Command{
		Short: "Create commit with a suitable emojiproviders based on the message of the comment using AI",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		Use:   "commit [commit massage]",
		Run: func(cmd *cobra.Command, args []string) {
			g, err := git.SetupGit()
			if err != nil {
				cmd.PrintErrln(err)
			}
			openAiToken := viper.GetString("open-ai-api-token")
			httpClient := &http.Client{}
			openAiClient := openai.NewClient(
				httpClient,
				openAiUrl,
				openAiToken,
				openai.Model35turbo,
			)
			emojiProvider := emojiproviders.NewChatGpt(openAiClient)
			aiGitEmoji := internal.NewAIGitEmoji(emojiProvider, g)
			commit := args[0]

			emojiCommit, hash, err := aiGitEmoji.Commit(cmd.Context(), commit, amend)
			if err != nil {
				cmd.PrintErrln(err)
			}

			cmd.Println(emojiCommit)
			cmd.Println(hash)

		},
	}

	flags := cmd.Flags()
	flags.BoolVar(&amend, "amend", false, "Replace the tip of the current branch "+
		"by creating a new commit. The recorded tree is prepared as usual "+
		"(including the effect of the -i and -o options and explicit pathspec), "+
		"and the message from the original commit is used as the starting point, "+
		"instead of an empty message, when no other message is specified from the command line via "+
		"options such as -m, -F, -c, etc. The new commit has the same parents and author "+
		"as the current one (the --reset-author option can countermand this).",
	)
	flags.StringVar(&openAiUrl, "open-ai-url", "https://api.openai.com", "open-ai url")
	flags.StringVar(&token, "open-ai-api-token", "", "username, facultative if you have config file")
	if err := viper.BindPFlag("open-ai-api-token", flags.Lookup("open-ai-api-token")); err != nil {
		fmt.Println(err)
	}

	return cmd
}
