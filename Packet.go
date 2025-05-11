package alphasign

//https://www.alpha-american.com/alpha-manuals/M-Protocol.pdf

import "fmt"

type Packet struct {
	SignAddress SignAddress
	SignType    SignTypeCode
	CommandCode CommandCode
	DataField   Bytes
}

func (p Packet) Bytes() []byte {
	out := []byte{
		byte(SOH),
		byte(p.SignType),
	}
	out = append(out, []byte(fmt.Sprintf("%02x", p.SignAddress))...)
	out = append(out, byte(STX), byte(p.CommandCode))
	out = append(out, p.DataField.Bytes()...)
	out = append(out, byte(EOT))
	return out
}
