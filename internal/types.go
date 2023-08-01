package internal

type (
	// Command each golang app creation is a command
	Command struct {
		GoVersion       string `survey:"go-version"`
		Architecture    string `survey:"architecture"`
		OperatingSystem string `survey:"operating-system"`
		Port            string `survey:"system-port"`
		SubCommands     []string
	}
)
