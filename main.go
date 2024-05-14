package main

import (
	"fmt"
	"lowleveldesign/elevator_system"
	"lowleveldesign/parkingSpot"
	"lowleveldesign/tic_toc_game"
)

func main() {

	lldMap := make(map[int]string)
	lldMap[0] = "Tic Toc Toe Game"
	lldMap[1] = "Parking Spot"
	lldMap[2] = "Elevator System"
	fmt.Println("Enter value to show lld design:-")

	for key, value := range lldMap {
		fmt.Println(key, "	", value)
	}

	var val int
	fmt.Scanln(&val)

	switch val {
	case 0:
		fmt.Println(lldMap[0])
		tic_toc_game.PlayGame()
	case 1:
		fmt.Println(lldMap[1])
		parkingSpot.ParkingSystem()
	case 2:
		fmt.Println(lldMap[1])
		elevator_system.ElevatorSystem()
	default:
		fmt.Println("Invalid input")
	}

}
