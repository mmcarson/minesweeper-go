package main

import "minesweeper-go/board"
import "fmt"

func main() {
	b := board.NewBoard(10, 10, 10)
	fmt.Printf("\n%s\n\n", "Welcome to Minesweeper!")
	var outcome int
	for {
		b.PrintBoard()
		fmt.Println("Enter coordinates to reveal a square. Format 'X Y' where X is horizontal and Y is vertical and both are capitalized. ")
		var x, y string
		fmt.Scan(&x, &y)
		outcome = b.RevealSquare(x, y)
		if outcome > 0 {
			b.PrintBoard()
			break
		}
	}
	if outcome == 1 {
		fmt.Printf("\n\n%s\n", "You lose!")
	} else {
		fmt.Printf("\n\n%s\n", "You win!")
	}
}