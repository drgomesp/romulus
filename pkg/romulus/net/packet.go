package net

import (
	"bytes"
	"fmt"
)

type PacketID uint16

const NilPacketID = PacketID(0)

type ClientPacket interface {
	Decode(*PacketData) error
}

type ServerPacket interface {
	Encode(*PacketData) error
}

type PacketReader interface {
	ReadPacket() (PacketID, ClientPacket, error)
}

type PacketWriter interface {
	Send(ServerPacket) error
	SendRaw(data interface{})
	Flush() error
}

type PacketExchange interface {
	PacketReader
	PacketWriter
}

type PacketHandlerFunc func(ClientPacket, PacketWriter) error

type PacketDefinition struct {
	Name    string
	ID      PacketID
	Size    int
	Data    *bytes.Buffer
	Decoder ClientPacket
}

func (t *PacketDefinition) String() string {
	return fmt.Sprintf("0x%04x", t.ID)
}
