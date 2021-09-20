package main

import (
	"fmt"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var square *ebiten.Image

// var activeSquare *ebiten.Image

var board [19][19]int

var screenWidth = len(board) * 20
var screenHeight = len(board[0]) * 20

// var screenWidth = 420 - 13
// var screenHeight = 420 - 13
var n = 0

// type _piece struct {
// 	X, Y  int
// 	Owner int
// }

// type _board struct {
// 	Grid [19][19]*_piece
// }

// func Board_init() *_board {
// 	return &_board{}
// }

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
	}
	boardString := ""
	for i := 0; i < 19; i++ {
		// if i < 10 {
		boardString += strconv.Itoa(19 - i)
		// }
		for j := 0; j < 19; j++ {
			if j == 0 && i < 10 {
				boardString += "| " + pieces[board[i][j]]
			} else {
				boardString += " | " + pieces[board[i][j]]
			}
			// boardString += " | " + pieces[board[i][j]]
		}
		boardString += " |\n"
	}
	boardString += "    A   B   C   D   E   F   G   H   I   J   K   L   M   N   O   P   Q   R   S"
	return boardString
}

func update(screen *ebiten.Image) error {
	screen.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	ebitenutil.DebugPrint(screen, "Our first game in Ebiten!")

	if square == nil {
		square, _ = ebiten.NewImage(19, 19, ebiten.FilterNearest)
	}
	for i := float64(0); i < float64(screenWidth); i += 20 {
		for j := float64(0); j < float64(screenHeight); j += 20 {
			square.Fill(color.White)
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
