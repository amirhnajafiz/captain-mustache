package internal

type (
	CommandType int
	StubType    int
)

const (
	MainCommand CommandType = iota + 1
	StubCommand
)

const (
	DatabaseStub StubType = iota + 1
	VolumeStub
	NetworkStub
)
