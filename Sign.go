package alphasign

import (
	"bufio"
	"io"

	"go.bug.st/serial"
)

type Sign struct {
	SignAddress SignAddress
	SignType    SignTypeCode
	Serial      io.ReadWriter
}

func New(signAddress SignAddress, signType SignTypeCode, devicePath string, baudRate int) (*Sign, error) {
	s := &Sign{
		SignAddress: signAddress,
		SignType:    signType,
	}

	mode := &serial.Mode{
		BaudRate: baudRate,
	}
	port, err := serial.Open(devicePath, mode)
	if err != nil {
		return nil, err
	}
	s.Serial = port

	return s, nil
}

func (s *Sign) Send(command Command) error {
	packet := &Packet{
		SignAddress: s.SignAddress,
		SignType:    s.SignType,
		CommandCode: command.CommandCode(),
		DataField:   command,
	}
	_, err := s.Serial.Write(packet.Bytes())
	return err
}

func (s *Sign) SendAndReceive(command CommandWithResponse, receiver any) error {
	err := s.Send(command)
	if err != nil {
		return err
	}
	reader := bufio.NewReader(s.Serial)
	response, err := reader.ReadSlice(byte(STX))
	if err != nil {
		return err
	}
	return command.ParseResponse(response, receiver)
}
