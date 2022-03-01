package net

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type ProtocolID string

type ProtocolHandlerFunc func(PacketExchange) error

type Protocol struct {
	ID  ProtocolID
	Run ProtocolHandlerFunc
}

type protoRW struct {
	logger *zap.SugaredLogger

	conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	reg    *PacketRegistry

	clientPacket *PacketDefinition
}

func (p *protoRW) ReadPacket() (PacketID, ClientPacket, error) {
	for true {
		var (
			packetID PacketID
			length   int
		)

		_ = length

		err := binary.Read(p.reader, binary.LittleEndian, &packetID)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if packetID == 0 {
			break
		}

		length, found := p.reg.PacketSize(packetID)
		if !found {
			return NilPacketID, nil, fmt.Errorf(`unknown packet "0x%04x"`, packetID)
		}

		variableSize := length == -1

		if variableSize {
			var s uint16

			if err := binary.Read(p.reader, binary.LittleEndian, &s); err != nil {
				return NilPacketID, nil, err
			}

			length = int(s)
		}

		data := make([]byte, length)
		if _, err := p.reader.Read(data); err != nil && err != io.EOF {
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

	return NilPacketID, nil, errors.New("kurwa")
}

func (p *protoRW) Send(packet ServerPacket) error {
	def, data, err := p.reg.Encode(packet)
	if err != nil {
		return err
	}

	_, err = p.writer.Write(data.Bytes())
	if err != nil {
		return err
	}

	p.logger.Debugw(
		"→ packet sent",
		"id", fmt.Sprintf("0x%04x", def.ID),
		"name", def.Name,
		"length", data.Len(),
		//"packet", packet,
	)

	return p.conn.Close()
}

func (p *protoRW) Flush() error {
	return p.writer.Flush()
}
