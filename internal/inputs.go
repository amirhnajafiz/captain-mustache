package internal

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// getInputs returns the base user command
func getInputs() (*Command, error) {
	cmd := new(Command)

	if err := survey.Ask(dockerfileQuestions, cmd); err != nil {
		return nil, fmt.Errorf("failed to get inputs error=%w", err)
	}

	list, er := getStubs()
	if er != nil {
		return nil, fmt.Errorf("failed to get sub inputs error=%w", er)
	}

	cmd.SubCommands = list

	return cmd, nil
}

// getStubs will create sub commands
func getStubs() ([]string, error) {
	var (
		commands      = make([]string, 0)
		databaseFlag  string
		dummyVariable string

		dbMessage = "Do you want to have database?"
	)

	// db selection
	for len(dbQuestion.Options) > 0 {
		if err := survey.AskOne(&survey.Select{
			Message: dbMessage,
			Options: []string{"Yes", "No"},
			Default: "No",
		}, &databaseFlag); err != nil {
			return nil, fmt.Errorf("failed to get survey input error=%w", err)
		}

		if databaseFlag == "Yes" {
			if err := survey.AskOne(dbQuestion, &dummyVariable); err != nil {
				return nil, fmt.Errorf("failed to get survey input error=%w", err)
			}

			commands = append(commands, dummyVariable)
			dbQuestion.Options = remove(dbQuestion.Options, dummyVariable)
			dbMessage = "Do you want to have another database?"
		} else {
			break
		}
	}

	return commands, nil
}
