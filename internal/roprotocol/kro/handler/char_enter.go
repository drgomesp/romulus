package handler

import (
	"github.com/pkg/errors"

	"github.com/drgomesp/rhizom/internal/roprotocol/kro/packet"
	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

func HandleCharEnter(clientPacket net.ClientPacket, sender net.PacketWriter) error {
	p, ok := clientPacket.(*packet.CharEnter)
	if !ok {
		return errors.New("fuk!")
	}

	err := sender.SendRaw(p.AccountID)

	if err != nil {
		return err
	}

	chars := loadCharacters()

	err = sender.Send(&packet.CharSlotsInfo{
		NormalSlots:     9,
		PremiumSlots:    0,
		BillingSlots:    0,
		ProducibleSlots: 9,
		ValidSlots:      9,
	})
	if err != nil {
		return err
	}

	err = sender.Send(&packet.CharAcceptEnter{
		TotalSlotCount:   9,
		PremiumSlotStart: 9,
		PremiumSlotEnd:   9,
		Chars:            chars,
	})
	if err != nil {
		return err
	}

	err = sender.Send(&packet.CharBlockCharacter{})
	if err != nil {
		return err
	}

	err = sender.Send(&packet.CharSecondPasswordLogin{
		Seed:      0xdeadbeef,
		AccountID: p.AccountID,
		Result:    0,
	})

	if err != nil {
		return err
	}

	err = sender.Send(&packet.CharListNotify{
		Count: 1,
	})

	if err != nil {
		return err
	}
	return nil
}

func loadCharacters() []*packet.CharacterInfo {
	return []*packet.CharacterInfo{
		packet.NewCharacterInfo("daniel"),
	}
}
