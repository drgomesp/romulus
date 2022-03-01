package net

import (
	"bytes"
	"encoding/binary"
	"log"
)

type PacketData struct {
	*bytes.Buffer
	ID PacketID
}

func NewPacketData(id PacketID) *PacketData {
	return &PacketData{
		Buffer: bytes.NewBuffer(make([]byte, 0)),
		ID:     id,
	}
}

func NewDataFromBytes(id PacketID, data []byte) *PacketData {
	return &PacketData{
		bytes.NewBuffer(data),
		id,
	}
}

func (d *PacketData) Read(v interface{}) {
	if err := binary.Read(d.Buffer, binary.LittleEndian, v); err != nil {
		log.Fatal(err)
	}
}

func (d *PacketData) ReadString(len int, s *string) {
	b := make([]byte, len)
	if _, err := d.Buffer.Read(b); err != nil {
		log.Fatal(err)
	}

	*s = string(b)
}

func (d *PacketData) Write(v interface{}) error {
	return binary.Write(d.Buffer, binary.LittleEndian, v)
}

func (d *PacketData) WriteString(size int, s string) error {
	str := []byte(s)
	data := make([]byte, size)

	if len(str) < size {
		copy(data, str)
	} else {
		copy(data, str[:size])
	}

	return d.Write(data)
}

func (d *PacketData) Skip(n int) {
	d.Buffer.Write(make([]byte, n))
}
