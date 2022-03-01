package packet

import (
	"time"

	"github.com/drgomesp/rhizom/pkg/romulus/net"
)

type CharacterInfo struct {
	ID           uint32
	BaseExp      uint64
	Zeny         uint32
	JobExp       uint64
	JobLevel     uint32
	BodyState    uint32
	HealthState  uint32
	EffectState  uint32
	Virtue       uint32
	Honor        uint32
	StatusPoints uint16
	HP           uint64
	MaxHP        uint64
	SP           uint64
	MaxSP        uint64
	Speed        uint16
	Job          uint16
	Head         uint16
	Body         uint16
	Weapon       uint16
	Level        uint16
	SkillPoints  uint16
	HeadBottom   uint16
	Shield       uint16
	HeadTop      uint16
	HeadMid      uint16
	HairPalette  uint16
	BodyPalette  uint16
	Name         string
	Str          byte
	Agi          byte
	Vit          byte
	Int          byte
	Dex          byte
	Luk          byte
	Slot         byte
	HairColor2   byte
	Renamed      bool
	MapName      string
	DeleteDate   *time.Time
	Robe         uint32
	SlotAddon    uint32
	RenameAddon  uint32
	Sex          bool
}

func (r *CharacterInfo) Encode(p *net.PacketData) error {
	p.Grow(175)

	p.Write(uint32(r.ID))
	p.Write(uint64(r.BaseExp))
	p.Write(uint32(r.Zeny))
	p.Write(uint64(r.JobExp))
	p.Write(uint32(r.JobLevel))
	p.Write(uint32(r.BodyState))
	p.Write(uint32(r.HealthState))
	p.Write(uint32(r.EffectState))
	p.Write(uint32(r.Virtue))
	p.Write(uint32(r.Honor))
	p.Write(uint16(r.StatusPoints))
	p.Write(uint64(r.HP))
	p.Write(uint64(r.MaxHP))
	p.Write(uint64(r.SP))
	p.Write(uint64(r.MaxSP))
	p.Write(uint16(r.Speed))
	p.Write(uint16(r.Job))
	p.Write(uint16(r.Head))
	p.Write(uint16(r.Head))
	p.Write(uint16(r.Weapon))
	p.Write(uint16(r.Level))
	p.Write(uint16(r.SkillPoints))
	p.Write(uint16(r.HeadBottom))
	p.Write(uint16(r.Shield))
	p.Write(uint16(r.HeadTop))
	p.Write(uint16(r.HeadMid))
	p.Write(uint16(r.HairPalette))
	p.Write(uint16(r.BodyPalette))
	p.WriteString(24, r.Name)
	p.Write(byte(r.Str))
	p.Write(byte(r.Agi))
	p.Write(byte(r.Vit))
	p.Write(byte(r.Int))
	p.Write(byte(r.Dex))
	p.Write(byte(r.Luk))
	p.Write(byte(r.Slot))
	p.Write(byte(r.HairColor2))

	if r.Renamed {
		p.Write(uint16(1))
	} else {
		p.Write(uint16(0))
	}

	p.WriteString(16, r.MapName)

	if r.DeleteDate != nil {
		p.Write(uint32(r.DeleteDate.Unix()))
	} else {
		p.Write(uint32(0))
	}

	p.Write(uint32(r.Robe))
	p.Write(uint32(r.SlotAddon))
	p.Write(uint32(r.RenameAddon))

	if r.Sex {
		p.Write(byte(1))
	} else {
		p.Write(byte(0))
	}

	return nil
}

func NewCharacterInfo(name string) *CharacterInfo {
	return &CharacterInfo{
		ID:           150002,
		BaseExp:      0,
		Zeny:         0,
		JobExp:       0,
		JobLevel:     1,
		BodyState:    0,
		HealthState:  0,
		EffectState:  0,
		Virtue:       0,
		Honor:        0,
		Body:         0,
		StatusPoints: 0,
		HP:           40,
		MaxHP:        40,
		SP:           11,
		MaxSP:        0,
		Speed:        150,
		Job:          0,
		Head:         0,
		Weapon:       0,
		SkillPoints:  0,
		Level:        1,
		HeadBottom:   0,
		Shield:       0,
		HeadTop:      0,
		HeadMid:      0,
		HairPalette:  0,
		BodyPalette:  0,
		Name:         name,
		Str:          1,
		Agi:          1,
		Vit:          1,
		Int:          1,
		Dex:          1,
		Luk:          1,
		Slot:         1,
		HairColor2:   0,
		Renamed:      true,
		MapName:      "prontera.gat",
		DeleteDate:   nil,
		Robe:         0,
		SlotAddon:    0,
		RenameAddon:  0,
		Sex:          true,
	}
}
