package internal

import "github.com/AlecAivazis/survey/v2"

var (
	dockerfileQuestions = []*survey.Question{
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
		{
			Name: "system-port",
			Prompt: &survey.Input{
				Message: "What is your system HTTP port on?",
				Default: "80",
			},
			Validate: survey.Required,
		},
	}

	dbQuestion = &survey.Select{
		Message: "Choose a database:",
		Options: []string{"redis", "mysql", "mongodb", "postgresql"},
	}
)
