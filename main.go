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

var square *ebiten.Image

// var activeSquare *ebiten.Image

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

func update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	ebitenutil.DebugPrint(screen, "Our first game in Ebiten!")

	if square == nil {
		square, _ = ebiten.NewImage(int(w), int(h), ebiten.FilterNearest)
	}
	for i := float64(0); i < float64(screenWidth); i += w {
		for j := float64(0); j < float64(screenHeight); j += h {
			if board[int(i/w)][int(j/h)] == 1 {
				square.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
			} else if board[int(i/w)][int(j/h)] == 2 {
				square.Fill(color.NRGBA{0x00, 0x00, 0xff, 0xff})
			} else {
				square.Fill(color.White)
			}
			if board[int(i/w)][int(j/h)] == 1 && board[int(i/w)+1][int(j/h)] == 2 && board[int(i/w)][int(j/h)+1] == 2 && board[int(i/w)][int(j/h)-1] == 2 && board[int(i/w)-1][int(j/h)] == 2 {
				board[int(i/w)][int(j/h)] = 0
			}
			if board[int(i/w)][int(j/h)] == 2 && board[int(i/w)+1][int(j/h)] == 1 && board[int(i/w)][int(j/h)+1] == 1 && board[int(i/w)][int(j/h)-1] == 1 && board[int(i/w)-1][int(j/h)] == 1 {
				board[int(i/w)][int(j/h)] = 0
			}
			x, y := ebiten.CursorPosition()
			ebitenutil.DebugPrint(screen, fmt.Sprintf("X: %d, Y: %d", x, y))
			if x/int(w) == int(i/w) {
				if y/int(h) == int(j/h) {
					if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
						board[int(i/w)][int(j/h)] = playerTurn()
						PrintBoard()
					}
				}
			}
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(i, j)
			screen.DrawImage(square, opts)
		}
	}
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
