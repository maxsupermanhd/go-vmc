package net

import pk "github.com/maxsupermanhd/go-vmc/v764/net/packet"

type Writer interface {
	WritePacket(p pk.Packet) error
}

type Reader interface {
	ReadPacket() (pk.Packet, error)
}

type ReadWriter interface {
	Reader
	Writer
}
