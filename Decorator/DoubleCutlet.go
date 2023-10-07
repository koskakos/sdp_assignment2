package main

type DoubleCutlet struct {
	burger Burger
}

func (b *DoubleCutlet) getPrice() int {
	return b.burger.getPrice() + 2
}
