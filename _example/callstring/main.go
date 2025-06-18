package main

import (
	"os"

	alphasign "github.com/johnjones4/alpha-sign-communications-protocol"
)

func main() {
	sign, err := alphasign.New(alphasign.SignAddressBroadcast, alphasign.AllSigns, os.Args[1], 9600)
	if err != nil {
		panic(err)
	}
	err = sign.Send(&alphasign.WriteSpecialFunctionCommand{
		Label: alphasign.ClearOrSetMemoryConfig,
		Data:  nil,
	})
	if err != nil {
		panic(err)
	}
	err = sign.Send(&alphasign.WriteSpecialFunctionCommand{
		Label: alphasign.ClearOrSetMemoryConfig,
		Data: alphasign.MemoryConfiguration{
			FileLabel:                'A',
			FileType:                 alphasign.TextFile,
			KeyboardProtectionStatus: 'L',
		},
	})
	if err != nil {
		panic(err)
	}
	err = sign.Send(&alphasign.WriteSpecialFunctionCommand{
		Label: alphasign.ClearOrSetMemoryConfig,
		Data: alphasign.MemoryConfiguration{
			FileLabel:                'B',
			FileType:                 alphasign.StringFile,
			KeyboardProtectionStatus: 'L',
		},
	})
	if err != nil {
		panic(err)
	}
	err = sign.Send(alphasign.WriteTextCommand{
		FileLabel: 'A',
		Mode: &alphasign.TextMode{
			DisplayPosition: alphasign.Left,
			ModeCode:        alphasign.Rotate,
		},
		Message: append([]byte("test: "), 0x10, 'B'),
	})
	if err != nil {
		panic(err)
	}
	err = sign.Send(alphasign.WriteStringCommand{
		FileLabel: 'B',
		FileData:  append([]byte{0x15, 0x1C, 0x31}, []byte("Hello World File Test!")...),
	})
	if err != nil {
		panic(err)
	}

}
