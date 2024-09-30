package main

import "fmt"

type City struct {
	cityId   int
	cityName string
	state    string
	country  string
}

func (c *City) GetTheaters() {
	fmt.Println("Fetching theaters in city:", c.cityName)
}
