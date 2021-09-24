package main

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var square, hLine, vLine *ebiten.Image

// var line *ebiten.Image

var board [19][19]int

var w, h float64 = 30, 30

var screenWidth = len(board) * int(w)
var screenHeight = len(board[0]) * int(h)

// var screenWidth = 420 - 13
// var screenHeight = 420 - 13
var n = 0

var playerOneTurn = true

func PopulateNewBoard() {
	board = [19][19]int{
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}
}

func PrintBoard() {
	fmt.Println(BoardAsString())
}

func BoardAsString() string {
	pieces := map[int]string{
		0: "0",
		1: "1",
		2: "2",
	}
	boardString := ""
	for i := 0; i < 19; i++ {
		boardString += strconv.Itoa(19 - i)
		for j := 0; j < 19; j++ {
			if j == 0 && i < 10 {
				boardString += "| " + pieces[board[i][j]]
			} else {
				boardString += " | " + pieces[board[i][j]]
			}
		}
		boardString += " |\n"
	}
	boardString += "    A   B   C   D   E   F   G   H   I   J   K   L   M   N   O   P   Q   R   S"
	return boardString
}

func playerTurn() int {
	if playerOneTurn == true {
		playerOneTurn = false
		return 1
	} else {
		playerOneTurn = true
		return 2
	}
}
func inSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func surrounded(board [19][19]int, i float64, j float64) int {
	// var s []int
	// if inSlice(int(i), s) == false {
	// if surrounded
	if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] != 0 && board[int(i/w)][int(j/h)+1] != 0 && board[int(i/w)-1][int(j/h)] != 0 && board[int(i/w)][int(j/h)-1] != 0 {
		// if surrounded by friends, good
		if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] == 1 && board[int(i/w)][int(j/h)+1] == 1 && board[int(i/w)-1][int(j/h)] == 1 && board[int(i/w)][int(j/h)-1] == 1 {
			return 1
			// if surrounded by enemies, bad
		} else if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] == 2 && board[int(i/w)][int(j/h)+1] == 2 && board[int(i/w)-1][int(j/h)] == 2 && board[int(i/w)][int(j/h)-1] == 2 {
			return 0
		}
		// if surrounded but touching a friend,...
		// } else if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] == 1 { //|| board[int(i/w)][int(j/h)+1] == 1 || board[int(i/w)-1][int(j/h)] == 1 || board[int(i/w)][int(j/h)-1] == 1 {
		// 	fmt.Println("HI")
		// 	// s = append(s, int(i), int(j))
		// 	// print(s[0])
		// 	board[int(i/w)][int(j/h)] = 2
		// 	return surrounded(board, i+1, j)
		// 	// PrintBoard()

		// 	// return surrounded(board, i+1, j)

		// }
	}
	// }
	// if surrounded
	if board[int(i/w)][int(j/h)] == 2 && board[int(i/w)+1][int(j/h)] != 0 && board[int(i/w)][int(j/h)+1] != 0 && board[int(i/w)-1][int(j/h)] != 0 && board[int(i/w)][int(j/h)-1] != 0 {
		// if surrounded by friends, good
		if board[int(i/w)][int(j/h)] == 2 && board[int(i/w)+1][int(j/h)] == 2 && board[int(i/w)][int(j/h)+1] == 2 && board[int(i/w)-1][int(j/h)] == 2 && board[int(i/w)][int(j/h)-1] == 2 {
			return 2
			// if surrounded by enemies, bad
		} else if board[int(i/w)][int(j/h)] == 2 && board[int(i/w)+1][int(j/h)] == 1 && board[int(i/w)][int(j/h)+1] == 1 && board[int(i/w)-1][int(j/h)] == 1 && board[int(i/w)][int(j/h)-1] == 1 {
			return 0
		}
	}
	return 3
}

func update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	ebitenutil.DebugPrint(screen, "Our first game in Ebiten!")

	if square == nil {
		square, _ = ebiten.NewImage(int(w), int(h), ebiten.FilterNearest)
	}

	if hLine == nil {
		hLine, _ = ebiten.NewImage(int(w), 1, ebiten.FilterNearest)
	}

	if vLine == nil {
		vLine, _ = ebiten.NewImage(1, int(h), ebiten.FilterNearest)
	}

	for i := float64(0); i < float64(screenWidth); i += w {
		hLine.Fill(color.Black)
		vLine.Fill(color.Black)
		for j := float64(0); j < float64(screenHeight); j += h {
			if board[int(i/w)][int(j/h)] == 1 {
				square.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
			} else if board[int(i/w)][int(j/h)] == 2 {
				square.Fill(color.NRGBA{0x00, 0x00, 0xff, 0xff})
			} else {
				square.Fill(color.White)
			}
			if surrounded(board, i, j) == 0 {
				// if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] != 0 && board[int(i/w)][int(j/h)+1] != 0 && board[int(i/w)-1][int(j/h)] != 0 && board[int(i/w)][int(j/h)-1] != 0 {
				// 	if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] == 1 && board[int(i/w)][int(j/h)+1] == 1 && board[int(i/w)-1][int(j/h)] == 1 && board[int(i/w)][int(j/h)-1] == 1 {
				// 		board[int(i/w)][int(j/h)] = 1
				// 	} else if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] == 2 && board[int(i/w)][int(j/h)+1] == 2 && board[int(i/w)-1][int(j/h)] == 2 && board[int(i/w)][int(j/h)-1] == 2 {
				// 		board[int(i/w)][int(j/h)] = 0
				// 	}
				// }
				board[int(i/w)][int(j/h)] = 0
			} else if surrounded(board, i, j) == 1 {
				board[int(i/w)][int(j/h)] = 1
			} else if surrounded(board, i, j) == 2 {
				board[int(i/w)][int(j/h)] = 2
			} else if surrounded(board, i, j) == 3 {
				if surrounded(board, i+1, j) == 1 {
					board[int(i/w)][int(j/h)] = 1
				}
			}
			// if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] != 0 && board[int(i/w)][int(j/h)+1] != 0 && board[int(i/w)-1][int(j/h)] != 0 && board[int(i/w)][int(j/h)-1] != 0 {
			// 	if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] == 1 && board[int(i/w)][int(j/h)+1] == 1 && board[int(i/w)-1][int(j/h)] == 1 && board[int(i/w)][int(j/h)-1] == 1 {
			// 		board[int(i/w)][int(j/h)] = 1
			// 	} else if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] == 2 && board[int(i/w)][int(j/h)+1] == 2 && board[int(i/w)-1][int(j/h)] == 2 && board[int(i/w)][int(j/h)-1] == 2 {
			// 		board[int(i/w)][int(j/h)] = 0
			// 	}
			// }

			// if board[int(i/w)][int(j/h)] == 2 && board[int(i/w)+1][int(j/h)] != 0 && board[int(i/w)][int(j/h)+1] != 0 && board[int(i/w)-1][int(j/h)] != 0 && board[int(i/w)][int(j/h)-1] != 0 {
			// 	if board[int(i/w)][int(j/h)] == 2 && board[int(i/w)+1][int(j/h)] == 1 || board[int(i/w)][int(j/h)+1] == 1 || board[int(i/w)-1][int(j/h)] == 1 || board[int(i/w)][int(j/h)-1] == 1 {
			// 		board[int(i/w)][int(j/h)] = 0
			// 		PrintBoard()
			// 	}
			// }
			// if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] == 2 && board[int(i/w)][int(j/h)+1] == 2 && board[int(i/w)][int(j/h)-1] == 2 && board[int(i/w)-1][int(j/h)] == 2 {
			// 	board[int(i/w)][int(j/h)] = 0
			// }
			// if board[int(i/w)][int(j/h)] == 2 && board[int(i/w)+1][int(j/h)] == 1 && board[int(i/w)][int(j/h)+1] == 1 && board[int(i/w)][int(j/h)-1] == 1 && board[int(i/w)-1][int(j/h)] == 1 {
			// 	board[int(i/w)][int(j/h)] = 0
			// }
			x, y := ebiten.CursorPosition()
			ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %d, Y: %d", x, y))
			if x/int(w) == int(i/w) {
				if y/int(h) == int(j/h) {
					if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
						board[int(i/w)][int(j/h)] = playerTurn()
						// PrintBoard()
					}
				}
			}
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(i, j)

			hLineOpts := &ebiten.DrawImageOptions{}
			hLineOpts.GeoM.Translate(i, j+h/2)

			vLineOpts := &ebiten.DrawImageOptions{}
			vLineOpts.GeoM.Translate(i+w/2, j)

			screen.DrawImage(square, opts)
			screen.DrawImage(hLine, hLineOpts)
			screen.DrawImage(vLine, vLineOpts)
		}
	}
	PrintBoard()
	// square.Fill(color.White)

	// opts := &ebiten.DrawImageOptions{}
	// // opts.GeoM.Translate(304, 64)
	// screen.DrawImage(square, opts)

	// for i := float64(0); i < float64(screenWidth); i += 11 {
	// 	for j := float64(0); j < float64(screenHeight); j += 11 {
	// 		if int(i)%2 == 1 {
	// 			square.Fill(color.White)
	// 			if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
	// 				ebitenutil.DebugPrint(screen, "You're pressing the 'LEFT' mouse button.")
	// 				square.Fill(color.NRGBA{0x00, 0x00, 0xff, 0xff})
	// 			}
	// 		} else {
	// 			square.Fill(color.Black)
	// 		}

	// 		if int(j)%2 == 0 {
	// 			square.Fill(color.Black)
	// 		}
	// 		opts := &ebiten.DrawImageOptions{}
	// 		opts.GeoM.Translate(i, j)
	// 		screen.DrawImage(square, opts)
	// 	}
	// 	// fmt.Println(ebiten.CursorPosition())
	// }

	// if activeSquare == nil {
	// 	activeSquare, _ = ebiten.NewImage(6, 6, ebiten.FilterNearest)
	// }
	// square.Fill(color.NRGBA{0x00, 0x00, 0xff, 0xff})
	// opts := &ebiten.DrawImageOptions{}
	// opts.GeoM.Translate(32, 21)
	// screen.DrawImage(square, opts)
	// if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
	// 	ebitenutil.DebugPrint(screen, "You're pressing the 'LEFT' mouse button.")
	// }
	return nil
}

func main() {
	PrintBoard()
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Hello world!"); err != nil {
		panic(err)
	}
}
