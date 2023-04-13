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

type Loadout interface {
	Team() string
	Operator() string
	OperatorIcon() string
	Primary() string
	PrimaryAttachments() Attachments
	Secondary() string
	SecondaryAttachments() Attachments
	Gadget() string
}

type AttackerLoadout struct {
	operator             Attacker
	primary              string
	primaryAttachments   Attachments
	secondary            string
	secondaryAttachments Attachments
	gadget               AttackerGadget
}

func (a AttackerLoadout) Team() string {
	return "Attacker"
}

func (a AttackerLoadout) Operator() string {
	return AttackerNames[a.operator]
}

func (a AttackerLoadout) OperatorIcon() string {
	return fmt.Sprintf("/assets/attackers/%d.png", a.operator)
}

func (a AttackerLoadout) Primary() string {
	return a.primary
}

func (a AttackerLoadout) PrimaryAttachments() Attachments {
	return a.primaryAttachments
}

func (a AttackerLoadout) Secondary() string {
	return a.secondary
}

func (a AttackerLoadout) SecondaryAttachments() Attachments {
	return a.secondaryAttachments
}

func (a AttackerLoadout) Gadget() string {
	return AttackerGadgetNames[a.gadget]
}

func GetRandomAttacker() AttackerLoadout {
	op := choose(Attackers)

	primary := choose(op.Primary)
	secondary := choose(op.Secondary)

	return AttackerLoadout{
		operator: op.Id,
		primary:  primary.Name,
		primaryAttachments: Attachments{
			Sight:       choose(primary.Sight),
			Barrel:      choose(primary.Barrel),
			Grip:        choose(primary.Grip),
			Underbarrel: choose(primary.Underbarrel),
		},
		secondary: secondary.Name,
		secondaryAttachments: Attachments{
			Sight:       choose(secondary.Sight),
			Barrel:      choose(secondary.Barrel),
			Grip:        choose(secondary.Grip),
			Underbarrel: choose(secondary.Underbarrel),
		},
		gadget: choose(op.Gadget),
	}
}

type DefenderLoadout struct {
	operator             Defender
	primary              string
	primaryAttachments   Attachments
	secondary            string
	secondaryAttachments Attachments
	gadget               DefenderGadget
}

func (a DefenderLoadout) Team() string {
	return "Defender"
}

func (d DefenderLoadout) Operator() string {
	return DefenderNames[d.operator]
}

func (d DefenderLoadout) OperatorIcon() string {
	return fmt.Sprintf("/assets/defenders/%d.png", d.operator)
}

func (d DefenderLoadout) Primary() string {
	return d.primary
}

func (d DefenderLoadout) PrimaryAttachments() Attachments {
	return d.primaryAttachments
}

func (d DefenderLoadout) Secondary() string {
	return d.secondary
}

func (d DefenderLoadout) SecondaryAttachments() Attachments {
	return d.secondaryAttachments
}

func (d DefenderLoadout) Gadget() string {
	return DefenderGadgetNames[d.gadget]
}

func GetRandomDefender() DefenderLoadout {
	op := choose(Defenders)

	primary := choose(op.Primary)
	secondary := choose(op.Secondary)

	return DefenderLoadout{
		operator: op.Id,
		primary:  primary.Name,
		primaryAttachments: Attachments{
			Sight:       choose(primary.Sight),
			Barrel:      choose(primary.Barrel),
			Grip:        choose(primary.Grip),
			Underbarrel: choose(primary.Underbarrel),
		},
		secondary: secondary.Name,
		secondaryAttachments: Attachments{
			Sight:       choose(secondary.Sight),
			Barrel:      choose(secondary.Barrel),
			Grip:        choose(secondary.Grip),
			Underbarrel: choose(secondary.Underbarrel),
		},
		gadget: choose(op.Gadget),
	}
}
