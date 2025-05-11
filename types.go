package alphasign

type ControlCharacter byte
type SignTypeCode byte
type SignAddress byte
type CommandCode byte
type FileLabel byte
type DisplayPosition byte
type StandardMode byte
type SpecialMode byte
type SpecialFunctionsLabel []byte
type FileType byte
type KeyboardProtectionStatus byte
type ColorStatus []byte
type SpecialTimeOfDay []byte

func (b ColorStatus) Bytes() []byte {
	return b
}

func (b SpecialTimeOfDay) Bytes() []byte {
	return b
}
