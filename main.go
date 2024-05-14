package main

import (
	"fmt"
	"lowleveldesign/carRental"
	"lowleveldesign/elevator_system"
	"lowleveldesign/loggingSystem"
	"lowleveldesign/parkingSpot"
	"lowleveldesign/tic_toc_game"
)

func main() {

	lldMap := make(map[int]string)
	lldMap[0] = "Tic Toc Toe Game"
	lldMap[1] = "Parking Spot"
	lldMap[2] = "Elevator System"
	lldMap[3] = "Car Rental"
	lldMap[4] = "Logging System"
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
	default:
		fmt.Println("Invalid input")
	}

}
