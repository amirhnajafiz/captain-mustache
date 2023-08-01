package internal

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// importBaseQuestions returns the base user inputs
func importBaseQuestions() (*Command, error) {
	cmd := new(BaseImports)

	if err := survey.Ask(dockerfileQuestions, cmd); err != nil {
		return nil, fmt.Errorf("failed to get inputs error=%w", err)
	}

	list, er := importSubCommands()
	if er != nil {
		return nil, fmt.Errorf("failed to get sub inputs error=%w", er)
	}

	return &Command{
		Imports:     cmd,
		SubCommands: list,
	}, nil
}

// importSubCommand will create sub commands
func importSubCommands() ([]*SubCommand, error) {
	var (
		commands     = make([]*SubCommand, 0)
		databaseFlag string
	)

	dbMessage := "Do you want to have database?"

	// db selection
	for {
		if len(dbQuestion.Options) == 0 {
			break
		}

		if err := survey.AskOne(&survey.Select{
			Message: dbMessage,
			Options: []string{"Yes", "No"},
			Default: "No",
		}, &databaseFlag); err != nil {
			return nil, fmt.Errorf("failed to get survey input error=%w", err)
		}

		if databaseFlag == "Yes" {
			var tmp string

			if err := survey.AskOne(dbQuestion, &tmp); err != nil {
				return nil, fmt.Errorf("failed to get survey input error=%w", err)
			}

			tmpCommand := &SubCommand{
				Param: tmp,
			}

			commands = append(commands, tmpCommand)
			dbQuestion.Options = remove(dbQuestion.Options, tmp)
			dbMessage = "Do you want to have another database?"
		} else {
			break
		}
	}

	return commands, nil
}
