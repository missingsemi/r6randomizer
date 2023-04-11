package main

import (
	"fmt"
	"math/rand"
)

type Attachments struct {
	Sight       Sight
	Barrel      Barrel
	Grip        Grip
	Underbarrel Underbarrel
}

func choose[T any](arr []T) T {
	if len(arr) == 0 {
		var res T
		return res
	}
	r := rand.Intn(len(arr))
	return arr[r]
}

const loadoutFmt string = "Operator: %s\nPrimary: %s\n* %s\n* %s\n* %s\n* %s\nSecondary: %s\n* %s\n* %s\n* %s\n* %s\nGadget: %s"

type AttackerLoadout struct {
	Operator             Attacker
	Primary              string
	PrimaryAttachments   Attachments
	Secondary            string
	SecondaryAttachments Attachments
	Gadget               AttackerGadget
}

func (a AttackerLoadout) String() string {
	return fmt.Sprintf(
		loadoutFmt,
		AttackerNames[a.Operator],
		a.Primary,
		SightNames[a.PrimaryAttachments.Sight],
		BarrelNames[a.PrimaryAttachments.Barrel],
		GripNames[a.PrimaryAttachments.Grip],
		UnderbarrelNames[a.PrimaryAttachments.Underbarrel],
		a.Secondary,
		SightNames[a.SecondaryAttachments.Sight],
		BarrelNames[a.SecondaryAttachments.Barrel],
		GripNames[a.SecondaryAttachments.Grip],
		UnderbarrelNames[a.SecondaryAttachments.Underbarrel],
		AttackerGadgetNames[a.Gadget],
	)
}

func GetRandomAttacker() AttackerLoadout {
	op := choose(Attackers)

	primary := choose(op.Primary)
	secondary := choose(op.Secondary)

	return AttackerLoadout{
		Operator: op.Id,
		Primary:  primary.Name,
		PrimaryAttachments: Attachments{
			Sight:       choose(primary.Sight),
			Barrel:      choose(primary.Barrel),
			Grip:        choose(primary.Grip),
			Underbarrel: choose(primary.Underbarrel),
		},
		Secondary: secondary.Name,
		SecondaryAttachments: Attachments{
			Sight:       choose(secondary.Sight),
			Barrel:      choose(secondary.Barrel),
			Grip:        choose(secondary.Grip),
			Underbarrel: choose(secondary.Underbarrel),
		},
		Gadget: choose(op.Gadget),
	}
}

type DefenderLoadout struct {
	Operator             Defender
	Primary              string
	PrimaryAttachments   Attachments
	Secondary            string
	SecondaryAttachments Attachments
	Gadget               DefenderGadget
}

func (a DefenderLoadout) String() string {
	return fmt.Sprintf(
		loadoutFmt,
		DefenderNames[a.Operator],
		a.Primary,
		SightNames[a.PrimaryAttachments.Sight],
		BarrelNames[a.PrimaryAttachments.Barrel],
		GripNames[a.PrimaryAttachments.Grip],
		UnderbarrelNames[a.PrimaryAttachments.Underbarrel],
		a.Secondary,
		SightNames[a.SecondaryAttachments.Sight],
		BarrelNames[a.SecondaryAttachments.Barrel],
		GripNames[a.SecondaryAttachments.Grip],
		UnderbarrelNames[a.SecondaryAttachments.Underbarrel],
		DefenderGadgetNames[a.Gadget],
	)
}

func GetRandomDefender() DefenderLoadout {
	op := choose(Defenders)

	primary := choose(op.Primary)
	secondary := choose(op.Secondary)

	return DefenderLoadout{
		Operator: op.Id,
		Primary:  primary.Name,
		PrimaryAttachments: Attachments{
			Sight:       choose(primary.Sight),
			Barrel:      choose(primary.Barrel),
			Grip:        choose(primary.Grip),
			Underbarrel: choose(primary.Underbarrel),
		},
		Secondary: secondary.Name,
		SecondaryAttachments: Attachments{
			Sight:       choose(secondary.Sight),
			Barrel:      choose(secondary.Barrel),
			Grip:        choose(secondary.Grip),
			Underbarrel: choose(secondary.Underbarrel),
		},
		Gadget: choose(op.Gadget),
	}
}
