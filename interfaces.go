package alphasign

type Bytes interface {
	Bytes() []byte
}

type Command interface {
	CommandCode() CommandCode
	Bytes
}

type CommandWithResponse interface {
	Command
	ParseResponse(response []byte, receiver any) error
}
