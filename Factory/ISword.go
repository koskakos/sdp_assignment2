package main

type ISword interface {
	setName(name string)
	setDamage(damage int)
	setRarity(rarity float32)
	getName() string
	getDamage() int
	getRarity() float32
}
