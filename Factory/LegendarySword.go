package main

type LegendarySword struct {
	Sword
}

func newLegendarySword() ISword {
	return &LegendarySword{
		Sword{
			name:   "Legendary Sword",
			damage: 10000,
			rarity: 0.000001,
		},
	}
}
