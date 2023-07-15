package internal

type (
	// SubCommand all network, stubs, volumes are subcommands
	SubCommand struct {
		Type  CommandType
		Stub  StubType
		Param string
	}

	// Command each golang app is a command
	Command struct {
		Version     string
		Image       string
		SubCommands []SubCommand
	}

	// Imports will be user base selects
	Imports struct {
	}
)
