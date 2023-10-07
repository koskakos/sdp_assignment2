package main

type Sword struct {
	name   string
	damage int
	rarity float32
}

func (s *Sword) setName(name string) {
	s.name = name
}

func (s *Sword) getName() string {
	return s.name
}

func (s *Sword) setDamage(damage int) {
	s.damage = damage
}

func (s *Sword) getDamage() int {
	return s.damage
}

func (s *Sword) setRarity(rarity float32) {
	s.rarity = rarity
}

func (s *Sword) getRarity() float32 {
	return s.rarity
}
