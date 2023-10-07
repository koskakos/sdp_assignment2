package main

import "fmt"

func main() {
	legendarySword, _ := getSword("Legendary Sword")
	commonSword, _ := getSword("Common Sword")

	printDetails(legendarySword)
	printDetails(commonSword)
}

func printDetails(sword ISword) {
	fmt.Printf("Sword name: %s \n Damage: %d \n Rarity: %f\n", sword.getName(), sword.getDamage(), sword.getRarity())
}
