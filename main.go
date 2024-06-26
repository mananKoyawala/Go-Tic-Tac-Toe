package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

var isPlayer1 = true

func main() {
	player_array := [3][3]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome, to Tic-Tac-Toe!")
	fmt.Println("It is two player game with O and X")
	fmt.Printf("\n")

	for {
		if isPlayer1 {
			color.Yellow("Player 1 : Choose the digit to set : O")
		} else {
			color.Cyan("Player 2 : Choose the digit to set : X")
		}

		PrintCanvas(&player_array)
		fmt.Println("Enter Number :")
		scanner.Scan()
		digit := scanner.Text()
		fmt.Printf("\n")
		if isPlayer1 {
			SetCanvas(digit, "O", &player_array, false)
		} else {
			SetCanvas(digit, "X", &player_array, true)
		}
		isGameOver(&player_array)
	}

}

func isGameOver(player_array *[3][3]string) {

	// Check row
	for i := 0; i < 3; i++ {
		if player_array[i][0] == player_array[i][1] && player_array[i][1] == player_array[i][2] {
			if player_array[i][1] == "O" {
				color.Green("Player 1 is Win!")
				PrintCanvas(player_array)
				os.Exit(0)
			}
			if player_array[i][1] == "X" {
				color.Green("Player 2 is Win!")
				PrintCanvas(player_array)
				os.Exit(0)
			}
		}
	}

	// Check column
	for j := 0; j < 3; j++ {
		if player_array[0][j] == player_array[1][j] && player_array[1][j] == player_array[2][j] {
			if player_array[1][j] == "O" {
				color.Green("Player 1 is Win!")
				PrintCanvas(player_array)
				os.Exit(0)
			}
			if player_array[1][j] == "X" {
				color.Green("Player 2 is Win!")
				PrintCanvas(player_array)
				os.Exit(0)
			}
		}
	}

	// Check diagonal
	if (player_array[0][0] == player_array[1][1] && player_array[1][1] == player_array[2][2]) ||
		(player_array[0][2] == player_array[1][1] && player_array[1][1] == player_array[2][0]) {
		if player_array[1][1] != " " {
			if player_array[1][1] == "O" {
				color.Green("Player 1 is Win!")
				PrintCanvas(player_array)
				os.Exit(0)
			}
			if player_array[1][1] == "X" {
				color.Green("Player 2 is Win!")
				PrintCanvas(player_array)
				os.Exit(0)
			}

		}
	}

	// Checks the moves available or not
	count := 0
	for _, row := range player_array {
		for _, val := range row {
			if val == "X" || val == "O" {
				count++
			}
		}
	}

	if count >= 9 {
		color.Magenta("Game is Draw, No Player is Win!")
		os.Exit(0)
	}
}

func PrintCanvas(player_array *[3][3]string) {
	fmt.Printf("\n")
	fmt.Printf("%v | %v | %v\n", player_array[0][0], player_array[0][1], player_array[0][2])
	fmt.Printf("---------\n")
	fmt.Printf("%v | %v | %v\n", player_array[1][0], player_array[1][1], player_array[1][2])
	fmt.Printf("---------\n")
	fmt.Printf("%v | %v | %v\n", player_array[2][0], player_array[2][1], player_array[2][2])
	fmt.Printf("\n")
}

func SetCanvas(d string, sign string, player_array *[3][3]string, set bool) {
	isavailable := false
	for i, row := range player_array {
		for j, val := range row {
			if d == val {
				isavailable = true
				player_array[i][j] = sign
				isPlayer1 = set
				break
			}
		}
	}

	if !isavailable {
		color.Red("Please select available space\n\n")
	}
}
