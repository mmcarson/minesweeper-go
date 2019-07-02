package board

import "fmt"
import "math/rand"
// import "strings"

var ALPHABET = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

type Board struct {
	squares [][]*Square
	Width int
	Height int
}

func NewBoard (width int, height int, mineChance int) *Board {
	b := new(Board)
	b.Width = width
	b.Height = height
	b.squares = make([][]*Square, width)
	for i := range b.squares {
		b.squares[i] = make([]*Square, height)
		for j := range b.squares[i] {
			b.squares[i][j] = NewSquare()
		}
	}
	b.addMines(mineChance)
	return b
}

func (b Board) addMines (mineChance int) {
	for i := range b.squares {
		for j := range b.squares[i] {
			if rand.Int() % mineChance == 0 {
				// make a mine
				b.squares[i][j].IsMine = true
				if i > 0 {
					if j > 0 {
						b.squares[i-1][j-1].NeighboringMines++
					}
					b.squares[i-1][j].NeighboringMines++
					if j < b.Height - 1 {
						b.squares[i-1][j+1].NeighboringMines++
					}
				}
				if j > 0 {
					b.squares[i][j-1].NeighboringMines++
				}
				if j < b.Height - 1 {
					b.squares[i][j+1].NeighboringMines++
				}
				if i < b.Width - 1 {
					if j > 0 {
						b.squares[i+1][j-1].NeighboringMines++
					}
					b.squares[i+1][j].NeighboringMines++
					if j < b.Height - 1 {
						b.squares[i+1][j+1].NeighboringMines++
					}
				}
			}
		}
	}
}

func (b Board) PrintBoard () {
	for i := range b.squares {
		if i == 0 {
			fmt.Printf("   ")
			for j := range b.squares[i] {
				fmt.Printf("  %s  ", ALPHABET[j])
			}
			fmt.Printf("\n\n")
		}
		fmt.Printf("%s  ", ALPHABET[i])
		for j := range b.squares[i] {
			fmt.Printf("  %s  ", b.squares[i][j].Print())
		}
		fmt.Printf("\n\n")
	}
}

func (b Board) RevealSquare (xString string, yString string) int {

	alphaMap := make(map[string]int)
	for i := range ALPHABET {
		alphaMap[ALPHABET[i]] = i
	}

	// splits := strings.Split(coordinates, " ")
	// fmt.Printf("%v\n", splits)
	x := alphaMap[xString]
	y := alphaMap[yString]

	square := b.squares[y][x]

	square.IsSeen = true

	if square.IsMine {
		return 1
	} else {
		for i := range b.squares {
			for j := range b.squares[i] {
				if !(b.squares[i][j].IsMine || b.squares[i][j].IsSeen) {
					return 0
				}
			}
		}
		return 2
	}
}