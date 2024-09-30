package main

import "fmt"

type Theater struct {
	theaterId    int
	name         string
	location     string
	totalScreens int
}

func (t *Theater) GetShowsByMovie() {
	fmt.Println("Fetching shows by movie in theater:", t.name)
}

func (t *Theater) GetScreens() {
	fmt.Println("Fetching screens in theater:", t.name)
}
