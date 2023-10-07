package main

type ExtraCheese struct {
	burger Burger
}

func (b *ExtraCheese) getPrice() int {
	return b.burger.getPrice() + 1
}
