package internal

type (
	// SubCommand all network, stubs, volumes are subcommands
	SubCommand struct {
		Stub  StubType
		Param string
	}

	// Command each golang app is a command
	Command struct {
		Imports     *BaseImports
		SubCommands []*SubCommand
	}

	// BaseImports will be user base selects
	BaseImports struct {
		GoVersion       string `survey:"go-version"`
		Architecture    string `survey:"architecture"`
		OperatingSystem string `survey:"operating-system"`
		Port            string `survey:"system-port"`
	}
)
