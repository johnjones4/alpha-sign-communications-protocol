package alphasign

import "fmt"

type WriteSpecialFunctionCommand struct {
	Label SpecialFunctionsLabel
	Data  Bytes
}

func (c *WriteSpecialFunctionCommand) CommandCode() CommandCode {
	return WriteSpecialFunction
}

func (c *WriteSpecialFunctionCommand) Bytes() []byte {
	if c.Data == nil {
		return c.Label
	}
	return append(c.Label, c.Data.Bytes()...)
}

type MemoryConfiguration struct {
	FileLabel                FileLabel
	FileType                 FileType
	KeyboardProtectionStatus KeyboardProtectionStatus
	FileSize                 Bytes
	Suffix                   *Bytes
}

func (c MemoryConfiguration) Bytes() []byte {
	bytes := []byte{
		byte(c.FileLabel),
		byte(c.FileType),
		byte(c.KeyboardProtectionStatus),
	}
	if c.FileSize != nil {
		bytes = append(bytes, c.FileSize.Bytes()...)
	}
	if c.Suffix != nil {
		bytes = append(bytes, (*c.Suffix).Bytes()...)
	} else {
		bytes = append(bytes, []byte("0000")...)
	}
	return bytes
}

type DotsPictureSize struct {
	Rows uint8
	Cols uint8
}

func (c DotsPictureSize) Bytes() []byte {
	return []byte(fmt.Sprintf("%2X%2X", c.Rows, c.Cols))
}

type StringFileSize struct{}

func (fs StringFileSize) Bytes() []byte {
	return []byte("0000")
}

type TimeOfDay struct {
	Hour uint16
	Min  uint16
}

func (t TimeOfDay) Bytes() []byte {
	if t.Min%10 != 0 {
		panic("minute must be divisible by 10")
	}

	offset := t.Hour*6 + t.Min

	return []byte(fmt.Sprintf("%2X", offset))
}

type StartEndTime struct {
	Start TimeOfDay
	End   TimeOfDay
}

func (t StartEndTime) Bytes() []byte {
	return append(t.Start.Bytes(), t.End.Bytes()...)
}
