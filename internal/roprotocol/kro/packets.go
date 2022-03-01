package kro

import (
	"github.com/drgomesp/rhizom/internal/roprotocol/kro/packet"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

var ClientPackets = map[string]net.ClientPacket{
	"CA_LOGIN":            &packet.AccountLogin{},
	"CH_ENTER":            &packet.ClientCharEnter{},
	"CH_REQUEST_CHARLIST": &packet.ClientRequestCharList{},
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
