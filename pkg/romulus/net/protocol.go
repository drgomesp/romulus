package net

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"net"

	"go.uber.org/zap"
)

type ProtocolID string

type ProtocolHandlerFunc func(PacketExchange) error

var ErrPacketType = func(pid PacketID) error {
	return fmt.Errorf("invalid packet '%04x'", pid)
}

type Protocol struct {
	ID  ProtocolID
	Run ProtocolHandlerFunc
}

type protoRW struct {
	logger *zap.SugaredLogger

	conn   net.Conn
	stream *bufio.ReadWriter
	reg    *PacketRegistry

	clientPacket *PacketDefinition
}

func (p *protoRW) ReadPacket() (PacketID, ClientPacket, error) {
	var (
		packetID PacketID
		length   int
	)

	err := binary.Read(p.stream, binary.LittleEndian, &packetID)
	if err != nil && err != io.EOF {
		return NilPacketID, nil, err
	}
	if packetID == NilPacketID {
		return NilPacketID, nil, nil
	}

	length, found := p.reg.PacketSize(packetID)
	if !found {
		return NilPacketID, nil, fmt.Errorf(`unknown packet "0x%04x"`, packetID)
	}

	variableSize := length == -1

	if variableSize {
		var s uint16

		if err := binary.Read(p.stream, binary.LittleEndian, &s); err != nil {
			return NilPacketID, nil, err
		}

		length = int(s)
	}

	data := make([]byte, length)
	if _, err := p.stream.Read(data); err != nil && err != io.EOF {
		return NilPacketID, nil, err
	}

	def, clientPacket, err := p.reg.Decode(NewDataFromBytes(packetID, data))
	if err != nil {
		return NilPacketID, nil, err
	}

	p.logger.Debugw("← packet received",
		"id", def,
		"name", def.Name,
		"length", length,
	)

	return packetID, clientPacket, nil
}

func (p *protoRW) Send(packet ServerPacket) error {
	def, data, err := p.reg.Encode(packet)
	if err != nil {
		return err
	}

	_, err = p.stream.Write(data.Bytes())
	if err != nil {
		return err
	}

	p.logger.Debugw(
		"→ packet sent",
		"id", fmt.Sprintf("0x%04x", def.ID),
		"name", def.Name,
		"length", data.Len(),
	)

	return p.stream.Flush()
}

func (p *protoRW) SendRaw(data interface{}) error {
	err := binary.Write(p.stream, binary.LittleEndian, data)

	if err != nil {
		return err
	}

	p.logger.Debugw("→ packet sent", "data", data)

	return p.stream.Flush()
}

func (p *protoRW) Error(err error) error {
	if err := p.conn.Close(); err != nil {
		return err
	}

	return err
}
