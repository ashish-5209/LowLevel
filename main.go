package main

import (
	"fmt"
	"lowleveldesign/atmSystem"
	"lowleveldesign/carRental"
	"lowleveldesign/elevator_system"
	"lowleveldesign/loggingSystem"
	"lowleveldesign/lruCache"
	"lowleveldesign/parkingSpot"
	"lowleveldesign/snakeLadderGame"
	"lowleveldesign/tic_toc_game"
)

func main() {

	lldMap := make(map[int]string)
	lldMap[0] = "Tic Toc Toe Game"
	lldMap[1] = "Parking Spot"
	lldMap[2] = "Elevator System"
	lldMap[3] = "Car Rental"
	lldMap[4] = "Logging System"
	lldMap[5] = "Snake Ladder Game"
	lldMap[6] = "ATM System"
	lldMap[7] = "LRU Cache"
	fmt.Println("Enter value to show lld design:-")

	for key, value := range lldMap {
		fmt.Println(key, "	", value)
	}

	var val int
	fmt.Scanln(&val)

	switch val {
	case 0:
		fmt.Println(lldMap[0])
		tic_toc_game.App()
	case 1:
		fmt.Println(lldMap[1])
		parkingSpot.App()
	case 2:
		fmt.Println(lldMap[2])
		elevator_system.App()
	case 3:
		fmt.Println(lldMap[3])
		carRental.App()
	case 4:
		fmt.Println(lldMap[4])
		loggingSystem.App()
	case 5:
		fmt.Println(lldMap[5])
		snakeLadderGame.App()
	case 6:
		fmt.Println(lldMap[6])
		atmSystem.App()
	case 7:
		fmt.Println(lldMap[7])
		lruCache.App()
	default:
		fmt.Println("Invalid input")
	}

}
