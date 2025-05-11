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
	err = sign.Send(alphasign.WriteTextCommand{
		FileLabel: 'A',
		Message:   []byte("Jack Jones"),
	})
	if err != nil {
		panic(err)
	}
}
