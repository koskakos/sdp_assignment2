package main

import "fmt"

func main() {
	burger := &MacBurger{}
	fmt.Println(burger.getPrice())

	burgerWithDoubleCutlet := &DoubleCutlet{burger: burger}
	fmt.Println(burgerWithDoubleCutlet.getPrice())

	burgerWithDoubleCutletAndExtraCheese := &ExtraCheese{burger: burgerWithDoubleCutlet}
	fmt.Println(burgerWithDoubleCutletAndExtraCheese.getPrice())
}
