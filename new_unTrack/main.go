package main



import (
	//"github.com/therealflyingfrog/new_untrack/untrace/completer"

	"fmt"
	"os"
	"os/exec"
	"strings"
	prompt "github.com/c-bata/go-prompt"
)

func executor(t string) {
	t = strings.TrimSpace(t)
	if t == "" {
		return
	} else if t == "quit" || t == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	cmd := exec.Command("cmd", "/C", t)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
	return
}




var suggestions = []prompt.Suggest{
	// Command , Description
	{Text: "dir", Description:"Show files in directory"},
}

func completer(t prompt.Document) []prompt.Suggest {
	w := t.GetWordBeforeCursor()
	if w == "" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(suggestions, w, true)
}



func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("unTrace => "),
		prompt.OptionPrefixTextColor(prompt.Red),
		prompt.OptionPreviewSuggestionTextColor(prompt.Green),
		prompt.OptionSelectedSuggestionBGColor(prompt.Turquoise),
		prompt.OptionSuggestionBGColor(prompt.Blue),
	)
	p.Run()
}
