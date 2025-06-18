package alphasign

const (
	SOH ControlCharacter = 0x01
	STX ControlCharacter = 0x02
	EOT ControlCharacter = 0x04
	ESC ControlCharacter = 0x1B

	Off ControlCharacter = 0x30
	On  ControlCharacter = 0x31

	DoubleHeightControl   ControlCharacter = 0x05
	TrueDescendersControl ControlCharacter = 0x06
	CharacterFlashControl ControlCharacter = 0x07

	ExtendedCharacterSets ControlCharacter = 0x08

	Celsius    ControlCharacter = 0x1C
	Fahrenheit ControlCharacter = 0x1D

	NoHoldSpeed ControlCharacter = 0x09

	CallDate              ControlCharacter = 0x0B
	DateMMDDYYWithSlashes ControlCharacter = 0x30
	DateDDMMYYWithSlashes ControlCharacter = 0x31
	DateMMDDYYWithDashes  ControlCharacter = 0x32
	DateDDMMYYWithDashes  ControlCharacter = 0x33
	DateMMDDYYWithDots    ControlCharacter = 0x34
	DateDDMMYYWithDots    ControlCharacter = 0x35
	DateMMDDYYWithSpaces  ControlCharacter = 0x36
	DateDDMMYYWithSpaces  ControlCharacter = 0x37
	DateMMMDDYYYY         ControlCharacter = 0x38
	DateDayOfWeek         ControlCharacter = 0x39

	NewPage               ControlCharacter = 0x0C
	NewLine               ControlCharacter = 0x0D
	SpeedControl          ControlCharacter = 0x0F
	CallString            ControlCharacter = 0x10
	DisableWideCharacters ControlCharacter = 0x11
	EnableWideCharacters  ControlCharacter = 0x12
	CallTime              ControlCharacter = 0x13
	CallSmallDotsPicture  ControlCharacter = 0x14

	Speed1 ControlCharacter = 0x15
	Speed2 ControlCharacter = 0x16
	Speed3 ControlCharacter = 0x17
	Speed4 ControlCharacter = 0x18
	Speed5 ControlCharacter = 0x19
)

const (
	AllSigns SignTypeCode = 'Z'
)

const SignAddressBroadcast SignAddress = 0x00

const (
	WriteText                CommandCode = 'A'
	ReadText                 CommandCode = 'B'
	WriteSpecialFunction     CommandCode = 'E'
	ReadSpecialFunction      CommandCode = 'F'
	WriteString              CommandCode = 'G'
	ReadString               CommandCode = 'H'
	WriteSmallDotsPicture    CommandCode = 'I'
	ReadSmallDotsPicture     CommandCode = 'J'
	WriteRGBDotsPicture      CommandCode = 'K'
	ReadRGBDotsPicture       CommandCode = 'L'
	WriteLargeDotsPicture    CommandCode = 'M'
	ReadLargeDotsPicture     CommandCode = 'N'
	WriteAlphaVisionBulletin CommandCode = 'O'
	SetTimeoutMessage        CommandCode = 'T'
)

const (
	MiddleLine DisplayPosition = 0x20
	TopLine    DisplayPosition = 0x22
	BottomLine DisplayPosition = 0x26
	Fill       DisplayPosition = 0x30
	Left       DisplayPosition = 0x31
	Right      DisplayPosition = 0x32
)

const (
	Rotate           StandardMode = 0x61
	Hold             StandardMode = 0x62
	Flash            StandardMode = 0x63
	RollUp           StandardMode = 0x65
	RollDown         StandardMode = 0x66
	RollLeft         StandardMode = 0x67
	RollRight        StandardMode = 0x68
	WipeUp           StandardMode = 0x69
	WipeDown         StandardMode = 0x6A
	WipeLeft         StandardMode = 0x6B
	WipeRight        StandardMode = 0x6C
	Scroll           StandardMode = 0x6D
	Automode         StandardMode = 0x6F
	RollIn           StandardMode = 0x70
	RollOut          StandardMode = 0x71
	WipeIn           StandardMode = 0x72
	WipeOut          StandardMode = 0x73
	CompressedRotate StandardMode = 0x74
	Explode          StandardMode = 0x75
	Clock            StandardMode = 0x76
	Special          StandardMode = 0x6E
)

const (
	Twinkle     SpecialMode = 0x30
	Sparkle     SpecialMode = 0x31
	Snow        SpecialMode = 0x32
	Interlock   SpecialMode = 0x33
	Switch      SpecialMode = 0x34
	Slide       SpecialMode = 0x35
	Spray       SpecialMode = 0x36
	Starburst   SpecialMode = 0x37
	Welcome     SpecialMode = 0x38
	SlotMachine SpecialMode = 0x39
	NewsFlash   SpecialMode = 0x3A
	Trumpet     SpecialMode = 0x3B
	CycleColors SpecialMode = 0x43
)

var (
	SetTimeOfDay                         SpecialFunctionsLabel = []byte{0x20}
	EnableDisableSpeaker                 SpecialFunctionsLabel = []byte{0x21}
	ClearOrSetMemoryConfig               SpecialFunctionsLabel = []byte{0x24}
	SetDayOfWeek                         SpecialFunctionsLabel = []byte{0x26}
	SetTimeFormat                        SpecialFunctionsLabel = []byte{0x27}
	GenerateSpeakerTone                  SpecialFunctionsLabel = []byte{0x28}
	SetRunTimeTable                      SpecialFunctionsLabel = []byte{0x29}
	DisplayTextAtXYPos                   SpecialFunctionsLabel = []byte{0x2B}
	SoftReset                            SpecialFunctionsLabel = []byte{0x2C}
	SetRunSequence                       SpecialFunctionsLabel = []byte{0x2E}
	SetDimmingRegister                   SpecialFunctionsLabel = []byte{0x2F}
	SetRunDayTable                       SpecialFunctionsLabel = []byte{0x32}
	ClearSerialErrorStatusRegister       SpecialFunctionsLabel = []byte{0x34}
	SetCounter                           SpecialFunctionsLabel = []byte{0x35}
	SetSerialAddress                     SpecialFunctionsLabel = []byte{0x37}
	SetLargeDotsMemoryConfiguration      SpecialFunctionsLabel = []byte{0x38}
	AppendToLargeDotsMemoryConfiguration SpecialFunctionsLabel = []byte{0x39}
	SetRunFileTimes                      SpecialFunctionsLabel = []byte{0x3A}
	SetDate                              SpecialFunctionsLabel = []byte{0x3B}
	ProgramCustomCharacterSet            SpecialFunctionsLabel = []byte{0x3C}
	SetAutomodeTable                     SpecialFunctionsLabel = []byte{0x3E}
	SetDimmingControlRegister            SpecialFunctionsLabel = []byte{0x40}
	SetColorCorrectionTable              SpecialFunctionsLabel = []byte{0x43, 0x33}
	SetCustomerColorCorrectionTable      SpecialFunctionsLabel = []byte{0x43, 0x58}
	SetTemperatureOffset                 SpecialFunctionsLabel = []byte{0x54}
	SetUnitColumnsAndRows                SpecialFunctionsLabel = []byte{0x55, 0x31}
	SetUnitRunMode                       SpecialFunctionsLabel = []byte{0x55, 0x32}
	SetUnitSerialAddress                 SpecialFunctionsLabel = []byte{0x55, 0x33}
	SetUnitSerialData                    SpecialFunctionsLabel = []byte{0x55, 0x34}
	SetUnitConfiguration                 SpecialFunctionsLabel = []byte{0x55, 0x35}
	SetUnitInternalNetwork               SpecialFunctionsLabel = []byte{0x55, 0x37}
	SetUnitSlaveDevice                   SpecialFunctionsLabel = []byte{0x55, 0x39}
	WriteUnitRegister                    SpecialFunctionsLabel = []byte{0x55, 0x4E}
	EnableDisableACKNAKResponse          SpecialFunctionsLabel = []byte{0x73}
)

var (
	TextFile        FileType = 'A'
	StringFile      FileType = 'B'
	DotsPictureFile FileType = 'C'
)

var (
	Monochrome ColorStatus = []byte("1000")
	ThreeColor ColorStatus = []byte("2000")
	EightColor ColorStatus = []byte("4000")
)

var (
	AllDay SpecialTimeOfDay = []byte("FD")
	Never  SpecialTimeOfDay = []byte("FE")
	Always SpecialTimeOfDay = []byte("FF")
)
