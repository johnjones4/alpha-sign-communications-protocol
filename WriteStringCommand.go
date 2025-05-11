package alphasign

type WriteStringCommand struct {
	FileLabel FileLabel
	FileData  []byte
}

func (c WriteStringCommand) CommandCode() CommandCode {
	return WriteString
}

func (c WriteStringCommand) Bytes() []byte {
	out := []byte{byte(c.FileLabel)}
	out = append(out, c.FileData...)
	return out
}
