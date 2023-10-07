package main

import "fmt"

func getSword(swordName string) (ISword, error) {
	if swordName == "Legendary Sword" {
		return newLegendarySword(), nil
	}
	if swordName == "Common Sword" {
		return newCommonSword(), nil
	}
	return nil, fmt.Errorf("Wrong sword type")
}
