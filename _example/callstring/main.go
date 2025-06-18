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
		Data: []alphasign.Bytes{
			alphasign.MemoryConfiguration{
				FileLabel:                'A',
				FileType:                 alphasign.TextFile,
				KeyboardProtectionStatus: 'U',
				FileSize:                 alphasign.FileSize(1024),
			},
			alphasign.MemoryConfiguration{
				FileLabel:                '1',
				FileType:                 alphasign.StringFile,
				KeyboardProtectionStatus: 'L',
				FileSize:                 alphasign.FileSize(1024),
			},
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
		Message: append([]byte("test"), 0x10, '1'),
	})
	if err != nil {
		panic(err)
	}

	err = sign.Send(alphasign.WriteStringCommand{
		FileLabel: '1',
		FileData:  []byte("Hello World File Test!"),
	})
	if err != nil {
		panic(err)
	}

}
