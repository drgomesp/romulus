package kro

import (
	"github.com/drgomesp/rhizom/internal/roprotocol/kro/packet"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

const (
	PacketClientPing                = net.PacketID(0x0187)
	PacketClientAccountLogin        = net.PacketID(0x0064)
	PacketClientCharEnter           = net.PacketID(0x0065)
	PacketClientCharSelectChar      = net.PacketID(0x0066)
	PacketClientCharRequestCharList = net.PacketID(0x09a1)

	PacketServerAccountAcceptLogin      = net.PacketID(0x0ac4)
	PacketServerCharAcceptEnter         = net.PacketID(0x006b)
	PacketServerCharSlotsInfo           = net.PacketID(0x082d)
	PacketServerCharSecondPasswordLogin = net.PacketID(0x08b9)
	PacketServerCharBlockCharacter      = net.PacketID(0x020d)
	PacketServerCharListNotify          = net.PacketID(0x09a0)
)

var ClientPackets = map[string]net.ClientPacket{
	"PING":                &packet.Ping{},
	"CA_LOGIN":            &packet.AccountLogin{},
	"CH_ENTER":            &packet.CharEnter{},
	"CH_REQUEST_CHARLIST": &packet.CharRequestCharList{},
	"CH_SELECT_CHAR":      &packet.CharSelectChar{},
}

var ServerPackets = map[string]net.ServerPacket{
	"AC_ACCEPT_LOGIN3":       &packet.AccountAcceptLogin{},
	"HC_ACCEPT_ENTER":        &packet.CharAcceptEnter{},
	"HC_BLOCK_CHARACTER":     &packet.CharBlockCharacter{},
	"HC_SLOT_INFO":           &packet.CharSlotsInfo{},
	"HC_SECOND_PASSWD_LOGIN": &packet.CharSecondPasswordLogin{},
	"HC_CHARLIST_NOTIFY":     &packet.CharListNotify{},
}
