package main

type Underbarrel int

func (u Underbarrel) String() string {
	return UnderbarrelNames[u]
}

const (
	NoLaser = Underbarrel(iota) + 1
	Laser
)

var UnderbarrelNames map[Underbarrel]string = map[Underbarrel]string{
	NoLaser: "No Laser",
	Laser:   "Laser",
}

type Barrel int

func (b Barrel) String() string {
	return BarrelNames[b]
}

const (
	NoBarrel = Barrel(iota) + 1
	FlashHider
	Compensator
	MuzzleBreak
	Suppressor
	Extended
)

var BarrelNames map[Barrel]string = map[Barrel]string{
	NoBarrel:    "No Barrel",
	FlashHider:  "Flash Hider",
	Compensator: "Compensator",
	MuzzleBreak: "Muzzle Break",
	Suppressor:  "Supressor",
	Extended:    "Extended",
}

type Grip int

func (g Grip) String() string {
	return GripNames[g]
}

const (
	NoGrip = Grip(iota) + 1
	VerticalGrip
	AngledGrip
)

var GripNames map[Grip]string = map[Grip]string{
	NoGrip:       "No Grip",
	VerticalGrip: "Vertical Grip",
	AngledGrip:   "Angled Grid",
}

type Sight int

func (s Sight) String() string {
	return SightNames[s]
}

const (
	NoSight = Sight(iota) + 1
	ReflexA
	ReflexB
	ReflexC
	ReflexD
	RedDotA
	RedDotB
	RedDotC
	HoloA
	HoloB
	HoloC
	HoloD

	OneFiveX
	TwoX
	AcogA
	AcogB
	ThreeX
)

var SightNames map[Sight]string = map[Sight]string{
	NoSight:  "No Sight",
	ReflexA:  "Reflex A",
	ReflexB:  "Reflex B",
	ReflexC:  "Reflex C",
	ReflexD:  "Reflex D",
	RedDotA:  "Red Dot A",
	RedDotB:  "Red Dot B",
	RedDotC:  "Red Dot C",
	HoloA:    "Holo A",
	HoloB:    "Holo B",
	HoloC:    "Holo C",
	HoloD:    "Holo D",
	OneFiveX: "SCOPE 1.5x",
	TwoX:     "SCOPE 2.0x",
	AcogA:    "SCOPE 2.5x A",
	AcogB:    "SCOPE 2.5x B",
	ThreeX:   "SCOPE 3.0x",
}
