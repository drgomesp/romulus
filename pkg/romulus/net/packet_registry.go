package net

import (
	"encoding/binary"
	"fmt"
	"reflect"
)

type PacketRegistry struct {
	ClientPackets map[string]ClientPacket
	ServerPackets map[string]ServerPacket

	Incoming map[PacketID]*PacketDefinition
	Outgoing map[reflect.Type]*PacketDefinition
}

func NewPacketRegistry(
	version uint32,
	clientPackets map[string]ClientPacket,
	serverPackets map[string]ServerPacket,
) *PacketRegistry {
	db := &PacketRegistry{
		ClientPackets: clientPackets,
		ServerPackets: serverPackets,

		Incoming: make(map[PacketID]*PacketDefinition),
		Outgoing: make(map[reflect.Type]*PacketDefinition),
	}

	pv, found := PacketVersions[int(version)]
	if !found {
		return nil
	}

	for id, d := range pv.Packets {
		db.RegisterPacket(d.Packet, id, d.Size)
	}

	return db
}

func (d *PacketRegistry) RegisterPacket(name string, id PacketID, size int) {
	var (
		decoder ClientPacket
		encoder ServerPacket
	)

	decoder, isDecoder := d.ClientPackets[name]
	if !isDecoder {
		decoder = nil
	}

	encoder, isEncoder := d.ServerPackets[name]
	if !isEncoder {
		encoder = nil
	}

	def := &PacketDefinition{
		Name:    name,
		ID:      id,
		Size:    size,
		Decoder: decoder,
	}

	d.Incoming[id] = def

	if isEncoder {
		d.Outgoing[reflect.TypeOf(encoder).Elem()] = def
	}
}

func (d *PacketRegistry) Encode(packet ServerPacket) (*PacketDefinition, *PacketData, error) {
	typ := reflect.TypeOf(packet).Elem()
	def, ok := d.Outgoing[typ]
	if !ok {
		return nil, nil, fmt.Errorf(`unknown server packet "0x%04x"`, typ)
	}

	data := NewPacketData(def.ID)
	if err := data.Write(uint16(def.ID)); err != nil {
		return nil, nil, err
	}

	if def.Size == -1 {
		_ = data.Write(uint16(0))
	}

	err := packet.Encode(data)
	if err != nil {
		return nil, nil, err
	}

	if def.Size == -1 {
		dataBytes := data.Bytes()

		binary.LittleEndian.PutUint16(dataBytes[2:4], uint16(data.Len()))
	}

	return def, data, nil
}

func (d *PacketRegistry) Decode(data *PacketData) (*PacketDefinition, ClientPacket, error) {
	def, ok := d.Incoming[data.ID]
	if !ok {
		return nil, nil, fmt.Errorf("unknown client packet 0x%04x", data.ID)
	}

	typ := reflect.TypeOf(def.Decoder).Elem()
	packet := reflect.New(typ).Interface().(ClientPacket)

	err := packet.Decode(data)
	if err != nil {
		return nil, nil, err
	}

	return def, packet, nil
}

func (d *PacketRegistry) PacketSize(packetID PacketID) (int, bool) {
	def, ok := d.Incoming[packetID]

	if !ok {
		return 0, false
	}

	return def.Size, true
}