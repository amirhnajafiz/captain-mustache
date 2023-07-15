package internal

import "github.com/AlecAivazis/survey/v2"

var questions = []*survey.Question{
	{
		Name: "go-version",
		Prompt: &survey.Select{
			Message: "Choose your Go version:",
			Options: []string{"1.16", "1.17", "1.18", "1.19", "1.20"},
			Default: "1.20",
		},
	},
	{
		Name: "architecture",
		Prompt: &survey.Select{
			Message: "Choose your Go architecture:",
			Options: []string{"arm64", "amd64"},
			Default: "arm64",
		},
	},
	{
		Name: "operating-system",
		Prompt: &survey.Select{
			Message: "Choose your Go operating system:",
			Options: []string{"linux", "darwin", "windows"},
			Default: "linux",
		},
	},
}

func importBaseQuestions() (*BaseImports, error) {
	tmp := new(BaseImports)

	if err := survey.Ask(questions, tmp); err != nil {
		return nil, err
	}

	return tmp, nil
}
