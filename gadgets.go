package main

type AttackerGadget int

const (
	StunGrenade = AttackerGadget(iota)
	FragGrenade
	SmokeGrenade
	ImpactEMP
	BreachingCharge
	HardBreachDevice
	Claymore
)

var AttackerGadgetNames map[AttackerGadget]string = map[AttackerGadget]string{
	StunGrenade:      "Stun Grenade",
	FragGrenade:      "Frag Grenade",
	SmokeGrenade:     "Smoke Grenade",
	ImpactEMP:        "Impact EMP",
	BreachingCharge:  "Breaching Charge",
	HardBreachDevice: "Hard Breach Device",
	Claymore:         "Claymore",
}

type DefenderGadget int

const (
	BarbedWire = DefenderGadget(iota)
	BulletproofCamera
	ProximityAlarm
	ImpactGrenade
	NitroCell
	DeployableShield
)

var DefenderGadgetNames map[DefenderGadget]string = map[DefenderGadget]string{
	BarbedWire:        "Barbed Wire",
	BulletproofCamera: "Bulletproof Camera",
	ProximityAlarm:    "Proximity Alarm",
	ImpactGrenade:     "Impact Grenade",
	NitroCell:         "Nitro Cell",
	DeployableShield:  "Deployable Shield",
}
