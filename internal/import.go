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

// importBaseQuestions returns the base user inputs
func importBaseQuestions() (*BaseImports, error) {
	tmp := new(BaseImports)

	if err := survey.Ask(questions, tmp); err != nil {
		return nil, err
	}

	return tmp, nil
}

func importSubCommands() ([]*SubCommand, error) {
	var (
		commands    = make([]*SubCommand, 0)
		volumeFlag  string
		networkFlag string
	)

	// db selection

	// network selection
	if err := survey.AskOne(&survey.Select{
		Message: "Do you want to have network?",
		Options: []string{"Yes", "No"},
		Default: "No",
	}, &networkFlag); err != nil {
		return nil, err
	}

	// volume selection
	if err := survey.AskOne(&survey.Select{
		Message: "Do you want to have volume?",
		Options: []string{"Yes", "No"},
		Default: "No",
	}, &volumeFlag); err != nil {
		return nil, err
	}

	return nil, nil
}
