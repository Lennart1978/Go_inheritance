package main

import "fmt"

// Define a struct type called "position" to represent coordinates (x, y).
type position struct {
	x float64
	y float64
}

// Define a method "move" for the "position" struct that updates the coordinates.
func (p *position) move(x, y float64) {
	p.x += x
	p.y += y
}

// Define a method "teleport" for the "position" struct that sets the coordinates.
func (p *position) teleport(x, y float64) {
	p.x = x
	p.y = y
}

// Define a struct type called "player" that embeds the "position" struct and
// includes additional fields for name, health, and score.
type player struct {
	position
	name   string
	health int
	score  int
}

// Implement a "String" method for the "player" struct to format its information.
func (p player) String() string {
	return fmt.Sprintf("\n%s\nHealth:%d\nScore:%d\nPosition:%.2f, %.2f", p.name, p.health, p.score, p.x, p.y)
}

// Define a struct type called "enemy" that also embeds the "position" struct
// and includes an additional "health" field.
type enemy struct {
	position
	health int
}

// Implement a "specialmove" method for the "enemy" struct to modify its position.
func (e *enemy) specialmove(x, y float64) {
	e.x *= x
	e.y *= y
}

func main() {
	fmt.Println("Inheritance in Go: - Test program")
	for i := 0; i < 32; i++ {
		fmt.Printf("_")
	}

	// Create an instance of the "player" struct named "lennart".
	var lennart player

	// Create an instance of the "enemy" struct named "böse".
	var böse enemy

	// Initialize fields for the "lennart" and "böse" instances.
	lennart.name = "Lennart"
	lennart.health, böse.health = 100, 100

	// Move "lennart" and "böse" using the "move" method.
	lennart.move(12.5, 44.7)
	böse.move(33.0, 9.0)

	// Print the information of "lennart" and "böse".
	fmt.Println(lennart)
	fmt.Println(böse)

	// Use the "specialmove" method on "böse" and "teleport" method on "lennart".
	böse.specialmove(6.66, 6.66)
	lennart.teleport(64.0, 128.0)

	// Print the updated information of "lennart" and "böse".
	fmt.Println(lennart)
	fmt.Println(böse)
}
