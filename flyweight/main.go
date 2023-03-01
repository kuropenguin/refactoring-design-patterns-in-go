package main

import "fmt"

func main() {
	game := newGame()

	game.addCounterTerrorist(TerroristDressType)
	game.addCounterTerrorist(TerroristDressType)
	game.addCounterTerrorist(TerroristDressType)
	game.addCounterTerrorist(TerroristDressType)

	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)
	game.addCounterTerrorist(CounterTerroristDressType)

	dressFactoryInstance := getDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.dressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.getColor())
	}
}
