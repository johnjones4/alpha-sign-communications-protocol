package alphasign

type TextMode struct {
	DisplayPosition  DisplayPosition
	ModeCode         StandardMode
	SpecialSpecifier *SpecialMode
}

type WriteTextCommand struct {
	FileLabel FileLabel
	Mode      *TextMode
	Message   []byte
}

func (c WriteTextCommand) CommandCode() CommandCode {
	return WriteText
}

func (c WriteTextCommand) Bytes() []byte {
	out := []byte{byte(c.FileLabel)}
	if c.Mode != nil {
		out = append(out, byte(ESC), byte(c.Mode.DisplayPosition), byte(c.Mode.ModeCode))
		if c.Mode.ModeCode == Special {
			out = append(out, byte(*c.Mode.SpecialSpecifier))
		}
	}
	out = append(out, c.Message...)
	return out
}
