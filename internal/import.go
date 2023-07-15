package internal

import "github.com/AlecAivazis/survey/v2"

var questions = map[string]*survey.Question{
	"go-version": {
		Name: "go-version",
		Prompt: &survey.Select{
			Message: "Choose your Go version:",
			Options: []string{"1.16", "1.17", "1.18", "1.19", "1.20"},
			Default: "1.20",
		},
	},
	"arch": {
		Name: "architecture",
		Prompt: &survey.Select{
			Message: "Choose your Go architecture:",
			Options: []string{"arm64", "amd64"},
			Default: "arm64",
		},
	},
	"os": {
		Name: "operating-system",
		Prompt: &survey.Select{
			Message: "Choose your Go operating system:",
			Options: []string{"linux", "darwin", "windows"},
			Default: "linux",
		},
	},
}

func importQuestions() {

}
