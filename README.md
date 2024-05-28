# AI GIT EMOJI
*aigitemoji* is a program that adds a suitable emoji based on the message of the comment using AI

### How to install
1. Install go-lang 1.22.2 or higher - https://go.dev/doc/install
2. run `go install github.com/yurahaid/aigitemoji/cmd/aigitemoji@latest`
3. Export open-ai api key int env `export OPEN_AI_API_KEY={your_key}`
4. 
   1. Use AI git emoji `aigitemoji commit "Test commit"` or  `aigitemoji commit --amend "New commit message"`
   2. If you don't want to provide the entire commit text to the AI, 
      you can add a secondary argument that will provide the AI, 
      but your original text will be private and only used for the local commit. 
      `aigitemoji commit "{private commit text}" "{text to search for emoji that will be sent to AI API} `
5. Enjoy 🎉

![example](/docs/example.gif)