package kro

import (
	"github.com/drgomesp/rhizom/internal/roprotocol/kro/packet"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

const (
	PacketClientPing                    = net.PacketID(0x0187)
	PacketClientAccountLogin            = net.PacketID(0x0064) // [CA] client > account
	PacketClientCharEnter               = net.PacketID(0x0065) // [CH] client > char
	PacketClientCharRequestCharList     = net.PacketID(0x09a1) // [CH] client > char
	PacketServerAccountAcceptLogin      = net.PacketID(0x0ac4) // [AC] account > client
	PacketServerCharAcceptEnter         = net.PacketID(0x006b) // [HC] char > client
	PacketServerCharSlotsInfo           = net.PacketID(0x082d) // [HC] char > client
	PacketServerCharSecondPasswordLogin = net.PacketID(0x08b9) // [HC] char > client
	PacketServerCharBlockCharacter      = net.PacketID(0x020d) // [HC] char > client
	PacketServerCharListNotify          = net.PacketID(0x09a0) // [HC] char > client
)

var ClientPackets = map[string]net.ClientPacket{
	"CA_LOGIN":            &packet.AccountLogin{},
	"CH_ENTER":            &packet.CharEnter{},
	"CH_REQUEST_CHARLIST": &packet.CharRequestCharList{},
	"PING":                &packet.Ping{},
}

var ServerPackets = map[string]net.ServerPacket{
	"AC_ACCEPT_LOGIN3":       &packet.AccountAcceptLogin{},
	"HC_ACCEPT_ENTER":        &packet.CharAcceptEnter{},
	"HC_BLOCK_CHARACTER":     &packet.CharBlockCharacter{},
	"HC_SLOT_INFO":           &packet.CharSlotsInfo{},
	"HC_SECOND_PASSWD_LOGIN": &packet.CharSecondPasswordLogin{},
	"HC_CHARLIST_NOTIFY":     &packet.CharListNotify{},
}
