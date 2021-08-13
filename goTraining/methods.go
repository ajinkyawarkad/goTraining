
// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a player's batting average. The formula is hits / atBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

// Add imports.

import (
	"fmt"
)

// Declare a struct that represents a ball player.
// Include fields called name, atBats and hits.

type player struct{
	name string;
	atBats int;
	hits int;

}



// Declare a method that calculates the batting average for a player.
func ( p player ) average() int {
	return p.hits/p.atBats
	}

func main() {

	// Create a slice of players and populate each player
	// with field values.

	// Display the batting average for each player in the slice.

	

	// virat := player{"Kohli",100, 50}

	// ajinkya := player{"Rahne",80, 40}
	// rahul := player{"KL",90, 45}

	// hardik := player{"Pandya",80,40}

	playerSlice := [...]player{{name: "Virat", atBats: 5, hits: 10},
	{name: "Harrdik", atBats: 2, hits: 30},
	{name: "Ajinkya", atBats: 3, hits: 50},
	{name: "M S Dhoni", atBats: 4, hits: 60},
	{name: "Shikhar", atBats: 5, hits: 20}}

	// playerSlice = append(playerSlice , virat, ajinkya, rahul, hardik)

	
	fmt.Println(playerSlice)

	for _, v := range playerSlice {
		fmt.Println("Name : ",v.name)
		fmt.Println("average : ", v.average())
		}
		




}
