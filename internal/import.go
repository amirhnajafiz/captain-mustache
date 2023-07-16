package internal

import "github.com/AlecAivazis/survey/v2"

// importBaseQuestions returns the base user inputs
func importBaseQuestions() (*Command, error) {
	tmp := new(BaseImports)

	if err := survey.Ask(dockerfileQuestions, tmp); err != nil {
		return nil, err
	}

	list, er := importSubCommands()
	if er != nil {
		return nil, er
	}

	return &Command{
		Imports:     tmp,
		SubCommands: list,
	}, nil
}

// importSubCommand will create sub commands
func importSubCommands() ([]*SubCommand, error) {
	var (
		commands     = make([]*SubCommand, 0)
		databaseFlag string
		volumeFlag   string
		networkFlag  string
	)

	// db selection
	for {
		if err := survey.AskOne(&survey.Select{
			Message: "Do you want to have database?",
			Options: []string{"Yes", "No"},
			Default: "No",
		}, &databaseFlag); err != nil {
			return nil, err
		}

		if databaseFlag == "Yes" {
			var tmp string

			if err := survey.AskOne(dbQuestion, &tmp); err != nil {
				return nil, err
			}

			tmpCommand := &SubCommand{
				Type:  StubCommand,
				Stub:  DatabaseStub,
				Param: tmp,
			}

			commands = append(commands, tmpCommand)
		} else {
			break
		}
	}

	// network selection
	if err := survey.AskOne(&survey.Select{
		Message: "Do you want to have network?",
		Options: []string{"Yes", "No"},
		Default: "No",
	}, &networkFlag); err != nil {
		return nil, err
	}

	if networkFlag == "Yes" {
		tmpCommand := &SubCommand{
			Type: StubCommand,
			Stub: NetworkStub,
		}

		commands = append(commands, tmpCommand)
	}

	// volume selection
	if err := survey.AskOne(&survey.Select{
		Message: "Do you want to have volume?",
		Options: []string{"Yes", "No"},
		Default: "No",
	}, &volumeFlag); err != nil {
		return nil, err
	}

	if volumeFlag == "Yes" {
		tmpCommand := &SubCommand{
			Type: StubCommand,
			Stub: VolumeStub,
		}

		commands = append(commands, tmpCommand)
	}

	return commands, nil
}
