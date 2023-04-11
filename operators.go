package main

type Attacker int

const (
	Sledge = Attacker(iota)
	Thatcher
	Ash
	Thermite
	Twitch
	Montagne
	Glaz
	Fuze
	Blitz
	IQ
	Buck
	Blackbeard
	Capitao
	Hibana
	Jackal
	Ying
	Zofia
	Dokkaebi
	Lion
	Finka
	Maverick
	Nomad
	Gridlock
	Nokk
	Amaru
	Kali
	Iana
	Ace
	Zero
	Flores
	Osa
	Sens
	Grim
	Brava
)

var AttackerNames map[Attacker]string = map[Attacker]string{
	Sledge:     "Sledge",
	Thatcher:   "Thatcher",
	Ash:        "Ash",
	Thermite:   "Thermite",
	Twitch:     "Twitch",
	Montagne:   "Montagne",
	Glaz:       "Glaz",
	Fuze:       "Fuze",
	Blitz:      "Blitz",
	IQ:         "IQ",
	Buck:       "Buck",
	Blackbeard: "Blackbeard",
	Capitao:    "Capitao",
	Hibana:     "Hibana",
	Jackal:     "Jackal",
	Ying:       "Ying",
	Zofia:      "Zofia",
	Dokkaebi:   "Dokkaebi",
	Lion:       "Lion",
	Finka:      "Finka",
	Maverick:   "Maverick",
	Nomad:      "Nomad",
	Gridlock:   "Gridlock",
	Nokk:       "Nokk",
	Amaru:      "Amaru",
	Kali:       "Kali",
	Iana:       "Iana",
	Ace:        "Ace",
	Zero:       "Zero",
	Flores:     "Flores",
	Osa:        "Osa",
	Sens:       "Sens",
	Grim:       "Grim",
	Brava:      "Brava",
}

type Defender int

const (
	Smoke = Defender(iota)
	Mute
	Castle
	Pulse
	Rook
	Doc
	Kapkan
	Tachanka
	Jager
	Bandit
	Frost
	Valkyrie
	Caviera
	Echo
	Mira
	Lesion
	Ela
	Vigil
	Maestro
	Alibi
	Clash
	Kaid
	Mozzie
	Warden
	Goyo
	Wamai
	Oryx
	Melusi
	Aruni
	Thunderbird
	Thorn
	Azami
	Solis
)

var DefenderNames map[Defender]string = map[Defender]string{
	Smoke:       "Smoke",
	Mute:        "Mute",
	Castle:      "Castle",
	Pulse:       "Pulse",
	Rook:        "Rook",
	Doc:         "Doc",
	Kapkan:      "Kapkan",
	Tachanka:    "Tachanka",
	Jager:       "Jager",
	Bandit:      "Bandit",
	Frost:       "Frost",
	Valkyrie:    "Valkyrie",
	Caviera:     "Caviera",
	Echo:        "Echo",
	Mira:        "Mira",
	Lesion:      "Lesion",
	Ela:         "Ela",
	Vigil:       "Vigil",
	Maestro:     "Maestro",
	Alibi:       "Alibi",
	Clash:       "Clash",
	Kaid:        "Kaid",
	Mozzie:      "Mozzie",
	Warden:      "Warden",
	Goyo:        "Goyo",
	Wamai:       "Wamai",
	Oryx:        "Oryx",
	Melusi:      "Melusi",
	Aruni:       "Aruni",
	Thunderbird: "Thunderbird",
	Thorn:       "Thorn",
	Azami:       "Azami",
	Solis:       "Solis",
}

type GunInfo struct {
	Name        string        `json:"name"`
	Sight       []Sight       `json:"sight"`
	Barrel      []Barrel      `json:"barrel"`
	Grip        []Grip        `json:"grip"`
	Underbarrel []Underbarrel `json:"underbarrel"`
}

type AttackerInfo struct {
	Id        Attacker
	Primary   []GunInfo
	Secondary []GunInfo
	Gadget    []AttackerGadget
}

type DefenderInfo struct {
	Id        Defender
	Primary   []GunInfo
	Secondary []GunInfo
	Gadget    []DefenderGadget
}

var Attackers []AttackerInfo = []AttackerInfo{
	{
		Id: Sledge,
		Primary: []GunInfo{
			{
				Name:        "M590A1",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "L85A2",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P226 MK 25",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{FragGrenade, StunGrenade, ImpactEMP},
	},
	{
		Id: Thatcher,
		Primary: []GunInfo{
			{
				Name:        "M590A1",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "L85A2",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "AR33",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P226 MK 25",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{Claymore, BreachingCharge},
	},
	{
		Id: Ash,
		Primary: []GunInfo{
			{
				Name:        "G36C",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "R4C",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "5.7 USG",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M45 MEUSOC",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{Claymore, BreachingCharge},
	},
	{
		Id: Thermite,
		Primary: []GunInfo{
			{
				Name:        "M1014",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "556XI",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "5.7 USG",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M45 MEUSOC",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{SmokeGrenade, StunGrenade},
	},
	{
		Id: Twitch,
		Primary: []GunInfo{
			{
				Name:        "F2",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor, Extended},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "417",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB, ThreeX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SG-CQB",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Grip:        []Grip{NoGrip, VerticalGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P9",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "LFP586",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{Claymore, SmokeGrenade},
	},
	{
		Id: Montagne,
		Primary: []GunInfo{
			{
				Name: "Leroc Shield",
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P9",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "LFP586",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{ImpactEMP, HardBreachDevice, SmokeGrenade},
	},
	{
		Id: Glaz,
		Primary: []GunInfo{
			{
				Name:        "OTS-03",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "PMM",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name: "GONNE-6",
			},
			{
				Name:        "BEARING 9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{Claymore, SmokeGrenade, FragGrenade},
	},
	{
		Id: Fuze,
		Primary: []GunInfo{
			{
				Name: "BALLISTIC SHIELD",
			},
			{
				Name:        "6P41",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "AK-12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "PMM",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "GSH-18",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{SmokeGrenade, BreachingCharge, HardBreachDevice},
	},
	{
		Id: Blitz,
		Primary: []GunInfo{
			{
				Name: "G52-TACTICAL SHIELD",
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P12",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{SmokeGrenade, BreachingCharge},
	},
	{
		Id: IQ,
		Primary: []GunInfo{
			{
				Name:        "AUG A2",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "552 COMMANDO",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "G8A1",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P12",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{BreachingCharge, Claymore},
	},
	{
		Id: Buck,
		Primary: []GunInfo{
			{
				Name:        "C8-SFW",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "CAMRS",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB, ThreeX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "MK1 9mm",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name: "GONNE-6",
			},
		},
		Gadget: []AttackerGadget{StunGrenade, HardBreachDevice},
	},
	{
		Id: Blackbeard,
		Primary: []GunInfo{
			{
				Name:        "MK17 CQB",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SR-25",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB, ThreeX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "D-50",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{StunGrenade, Claymore, ImpactEMP},
	},
	{
		Id: Capitao,
		Primary: []GunInfo{
			{
				Name:        "PARA-308",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M249",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "PRB92",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name: "GONNE-6",
			},
		},
		Gadget: []AttackerGadget{HardBreachDevice, Claymore},
	},
	{
		Id: Hibana,
		Primary: []GunInfo{
			{
				Name:        "TYPE-89",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SUPERNOVA",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P229",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "BEARING-9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{StunGrenade, BreachingCharge},
	},
	{
		Id: Jackal,
		Primary: []GunInfo{
			{
				Name:        "C7E",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "PDW9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "ITA12L",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "USP40",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "ITA12S",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{SmokeGrenade, Claymore},
	},
	{
		Id: Ying,
		Primary: []GunInfo{
			{
				Name:        "T-95 LSW",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SIX12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "Q-929",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{SmokeGrenade, HardBreachDevice},
	},
	{
		Id: Zofia,
		Primary: []GunInfo{
			{
				Name:        "LMG-E",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M762",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "RG15",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{BreachingCharge, Claymore},
	},
	{
		Id: Dokkaebi,
		Primary: []GunInfo{
			{
				Name:        "MK 14 EBR",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB, ThreeX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "BOSG.12.2",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "SMG-12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "C75 Auto",
				Barrel:      []Barrel{NoBarrel, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name: "GONNE-6",
			},
		},
		Gadget: []AttackerGadget{SmokeGrenade, StunGrenade, ImpactEMP},
	},
	{
		Id: Lion,
		Primary: []GunInfo{
			{
				Name:        "V308",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "417",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB, ThreeX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SG-CQB",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Grip:        []Grip{NoGrip, VerticalGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P9",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "LFP586",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name: "GONNE-6",
			},
		},
		Gadget: []AttackerGadget{StunGrenade, Claymore, ImpactEMP},
	},
	{
		Id: Finka,
		Primary: []GunInfo{
			{
				Name:        "SPEAR.308",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "6P41",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB, ThreeX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "PMM",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "GSH-18",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name: "GONNE-6",
			},
		},
		Gadget: []AttackerGadget{SmokeGrenade, StunGrenade},
	},
	{
		Id: Maverick,
		Primary: []GunInfo{
			{
				Name:        "AR-15.50",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB, ThreeX},
				Barrel:      []Barrel{NoBarrel, Suppressor, MuzzleBreak},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M4",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "1911 TACOPS",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{Claymore, StunGrenade},
	},
	{
		Id: Nomad,
		Primary: []GunInfo{
			{
				Name:        "AK-74M",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "ARX200",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        ".44 MAG SEMI-AUTO",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "PRB92",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{StunGrenade, BreachingCharge},
	},
	{
		Id: Gridlock,
		Primary: []GunInfo{
			{
				Name:        "F90",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M249 SAW",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "SUPER SHORTY",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SDP 9mm",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name: "GONNE-6",
			},
		},
		Gadget: []AttackerGadget{SmokeGrenade, ImpactEMP, BreachingCharge},
	},
	{
		Id: Nokk,
		Primary: []GunInfo{
			{
				Name:        "FMG-9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SIX12 SD",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "5.7 USG",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "D-50",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{FragGrenade, ImpactEMP, HardBreachDevice},
	},
	{
		Id: Amaru,
		Primary: []GunInfo{
			{
				Name:        "G8A1",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SUPERNOVA",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "SMG-11",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "ITA12S",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name: "GONNE-6",
			},
		},
		Gadget: []AttackerGadget{HardBreachDevice, StunGrenade},
	},
	{
		Id: Kali,
		Primary: []GunInfo{
			{
				Name: "CSRX 300",
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "SPSMG9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "C75 Auto",
				Barrel:      []Barrel{NoBarrel, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "P226 MK 25",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{BreachingCharge, Claymore},
	},
	{
		Id: Iana,
		Primary: []GunInfo{
			{
				Name:        "ARX200",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "G36C",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "MK1 9mm",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{FragGrenade, SmokeGrenade},
	},
	{
		Id: Ace,
		Primary: []GunInfo{
			{
				Name:        "AK-12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M1014",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P9",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{BreachingCharge, Claymore},
	},
	{
		Id: Zero,
		Primary: []GunInfo{
			{
				Name:        "SC3000K",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "MP7",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "5.7 USG",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name: "GONNE-6",
			},
		},
		Gadget: []AttackerGadget{HardBreachDevice, Claymore},
	},
	{
		Id: Flores,
		Primary: []GunInfo{
			{
				Name:        "AR33",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SR-25",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB, ThreeX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "GSH-18",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{StunGrenade, Claymore},
	},
	{
		Id: Osa,
		Primary: []GunInfo{
			{
				Name:        "556XI",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "PDW9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "PMM",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{SmokeGrenade, Claymore, ImpactEMP},
	},
	{
		Id: Sens,
		Primary: []GunInfo{
			{
				Name:        "POF-9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "417",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB, ThreeX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "SDP 9mm",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name: "GONNE-6",
			},
		},
		Gadget: []AttackerGadget{HardBreachDevice, Claymore},
	},
	{
		Id: Grim,
		Primary: []GunInfo{
			{
				Name:        "552 COMMANDO",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SG-CQB",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Grip:        []Grip{NoGrip, VerticalGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P229",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{BreachingCharge, Claymore},
	},
	{
		Id: Brava,
		Primary: []GunInfo{
			{
				Name:        "PARA-308",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "CAMRS",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB, ThreeX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, AngledGrip, VerticalGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "SUPER SHORTY",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "USP40",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []AttackerGadget{BreachingCharge, Claymore},
	},
}

var Defenders []DefenderInfo = []DefenderInfo{
	{
		Id: Smoke,
		Primary: []GunInfo{
			{
				Name:        "FMG-9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M590A1",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P226 MK 25",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SMG-11",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BarbedWire, DeployableShield},
	},
	{
		Id: Mute,
		Primary: []GunInfo{
			{
				Name:        "MP5K",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M590A1",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P226 MK 25",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SMG-11",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{NitroCell, BulletproofCamera},
	},
	{
		Id: Castle,
		Primary: []GunInfo{
			{
				Name:        "UMP45",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M1014",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "5.7 USG",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SUPER SHORTY",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M45 MEUSOC",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ProximityAlarm, BulletproofCamera},
	},
	{
		Id: Pulse,
		Primary: []GunInfo{
			{
				Name:        "UMP45",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M1014",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "5.7 USG",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M45 MEUSOC",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{NitroCell, BarbedWire},
	},
	{
		Id: Doc,
		Primary: []GunInfo{
			{
				Name:        "SG-CQB",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Grip:        []Grip{NoGrip, VerticalGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "MP5",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "P90",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P9",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "LFP586",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "Bailiff 410",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BulletproofCamera, BarbedWire},
	},
	{
		Id: Rook,
		Primary: []GunInfo{
			{
				Name:        "SG-CQB",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Grip:        []Grip{NoGrip, VerticalGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "MP5",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "P90",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P9",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "LFP586",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ImpactGrenade, ProximityAlarm},
	},
	{
		Id: Kapkan,
		Primary: []GunInfo{
			{
				Name:        "9x19VSN",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SASG-12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "PMM",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "GSH-18",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ImpactGrenade, NitroCell},
	},
	{
		Id: Tachanka,
		Primary: []GunInfo{
			{
				Name:        "9x19VSN",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:  "DP27",
				Sight: []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, ReflexD},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "PMM",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "GSH-18",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "BEARING 9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BarbedWire, DeployableShield},
	},
	{
		Id: Jager,
		Primary: []GunInfo{
			{
				Name:        "M870",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "416-C CARBINE",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P12",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BulletproofCamera, BarbedWire},
	},
	{
		Id: Bandit,
		Primary: []GunInfo{
			{
				Name:        "MP7",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M870",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P12",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{NitroCell, BarbedWire},
	},
	{
		Id: Frost,
		Primary: []GunInfo{
			{
				Name:        "SUPER 90",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "9mm C1",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "MK1 9mm",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "ITA12S",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{DeployableShield, BulletproofCamera},
	},
	{
		Id: Valkyrie,
		Primary: []GunInfo{
			{
				Name:        "MPX",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SPAS-12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "D-50",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ImpactGrenade, NitroCell},
	},
	{
		Id: Caviera,
		Primary: []GunInfo{
			{
				Name:        "M12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SPAS-15",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "LUISON",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ImpactGrenade, ProximityAlarm},
	},
	{
		Id: Echo,
		Primary: []GunInfo{
			{
				Name:        "SUPERNOVA",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "MP5SD",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P229",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "BEARING 9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ImpactGrenade, DeployableShield},
	},
	{
		Id: Mira,
		Primary: []GunInfo{
			{
				Name:        "VECTOR .45 ACP",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "ITA12L",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "USP40",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "ITA12S",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ProximityAlarm, NitroCell},
	},
	{
		Id: Lesion,
		Primary: []GunInfo{
			{
				Name:        "T-5 SMG",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SIX12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "Q-929",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ImpactGrenade, BulletproofCamera},
	},
	{
		Id: Ela,
		Primary: []GunInfo{
			{
				Name:        "SCORPION EVO 3 AI",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "FO-12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "RG15",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BarbedWire, DeployableShield},
	},
	{
		Id: Vigil,
		Primary: []GunInfo{
			{
				Name:        "K1A",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "BOSG.12.2",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX, AcogA, AcogB},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "C75 Auto",
				Barrel:      []Barrel{NoBarrel, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SMG-12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BulletproofCamera, ImpactGrenade},
	},
	{
		Id: Maestro,
		Primary: []GunInfo{
			{
				Name:        "ALDA 5.56",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "ACS12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "Bailiff 410",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "KERATOS .357",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BarbedWire, ImpactGrenade},
	},
	{
		Id: Alibi,
		Primary: []GunInfo{
			{
				Name:        "Mx4 Storm",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "ACS12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "Bailiff 410",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "KERATOS .357",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{DeployableShield, ImpactGrenade},
	},
	{
		Id: Clash,
		Primary: []GunInfo{
			{
				Name: "CCE SHIELD",
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "SUPER SHORTY",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SPSMG9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "P-10C",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BarbedWire, ImpactGrenade},
	},
	{
		Id: Kaid,
		Primary: []GunInfo{
			{
				Name:        "AUG A3",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "TCSG12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        ".44 MAG SEMI-AUTO",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "LFP586",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BarbedWire, NitroCell},
	},
	{
		Id: Mozzie,
		Primary: []GunInfo{
			{
				Name:        "COMMANDO 9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "P10 RONI",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "SDP 9mm",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BarbedWire, NitroCell},
	},
	{
		Id: Warden,
		Primary: []GunInfo{
			{
				Name:        "M590A1",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "MPX",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P-10C",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SMG-12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{DeployableShield, NitroCell},
	},
	{
		Id: Goyo,
		Primary: []GunInfo{
			{
				Name:        "VECTOR .45 ACP",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "TCSG12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX, TwoX},
				Barrel:      []Barrel{NoBarrel, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "P229",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ProximityAlarm, NitroCell},
	},
	{
		Id: Wamai,
		Primary: []GunInfo{
			{
				Name:        "AUG A2",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "MP5K",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "KERATOS .357",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "P12",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ProximityAlarm, ImpactGrenade},
	},
	{
		Id: Oryx,
		Primary: []GunInfo{
			{
				Name:        "T-5 SMG",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SPAS-12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "Bailiff 410",
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "USP40",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ProximityAlarm, BarbedWire},
	},
	{
		Id: Melusi,
		Primary: []GunInfo{
			{
				Name:        "MP5",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SUPER 90",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "RG15",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BulletproofCamera, ImpactGrenade},
	},
	{
		Id: Aruni,
		Primary: []GunInfo{
			{
				Name:        "P10 RONI",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "MK 14 EBR",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "PRB92",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BulletproofCamera, BarbedWire},
	},
	{
		Id: Thunderbird,
		Primary: []GunInfo{
			{
				Name:        "SPEAR.308",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "SPAS-15",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "BEARING 9",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "Q-929",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ImpactGrenade, NitroCell},
	},
	{
		Id: Thorn,
		Primary: []GunInfo{
			{
				Name:        "UZK50GI",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, FlashHider, Compensator, MuzzleBreak, Suppressor, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "M870",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "1911 TACOPS",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "C75 Auto",
				Barrel:      []Barrel{NoBarrel, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{DeployableShield, BarbedWire},
	},
	{
		Id: Azami,
		Primary: []GunInfo{
			{
				Name:        "9x19VSN",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Grip:        []Grip{NoGrip, VerticalGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "ACS12",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "D-50",
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{BarbedWire, ImpactGrenade},
	},
	{
		Id: Solis,
		Primary: []GunInfo{
			{
				Name:        "P90",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC, OneFiveX},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Compensator, FlashHider, Extended},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
			{
				Name:        "ITA12L",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Secondary: []GunInfo{
			{
				Name:        "SMG-11",
				Sight:       []Sight{NoSight, RedDotA, RedDotB, RedDotC, HoloA, HoloB, HoloC, HoloD, ReflexA, ReflexB, ReflexC},
				Barrel:      []Barrel{NoBarrel, MuzzleBreak, Suppressor, Extended, Compensator, FlashHider},
				Grip:        []Grip{NoGrip, VerticalGrip, AngledGrip},
				Underbarrel: []Underbarrel{NoLaser, Laser},
			},
		},
		Gadget: []DefenderGadget{ImpactGrenade, BulletproofCamera},
	},
}
