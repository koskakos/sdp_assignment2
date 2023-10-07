package main

type CommonSword struct {
	Sword
}

func newCommonSword() ISword {
	return &CommonSword{
		Sword{
			name:   "Common Sword",
			damage: 1,
			rarity: 0.97,
		},
	}
}
