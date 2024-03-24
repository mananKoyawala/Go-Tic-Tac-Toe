# Tic-Tac-Toe Game in Go

Welcome to my Tic-Tac-Toe game implementation in Go! This project provides a simple command-line interface (CLI) implementation of the classic Tic-Tac-Toe game using the Go programming language.

## Features

- **Tic-Tac-Toe**: Play the classic Tic-Tac-Toe game against a friend.
- **CLI Interface**: Interact with the game through a user-friendly command-line interface.
- **Colorful Output**: Utilizes the `github.com/fatih/color` package to enhance the visual presentation of the game.

## Installation

1. Ensure you have Go installed on your machine. If not, you can download it [here](https://golang.org/dl/).

2. Clone this repository to your local machine:

```bash
   git clone https://github.com/mananKoyawala/Go-Tic-Tac-Toe.git
```

3. Navigate to the project directory:

```bash
   cd Go-Tic-Tac-Toe
```

4. Add the color package to your project by executing the following command:

```bash
   go get github.com/fatih/color
```

5. Run the game using the following command:

```bash
   go run main.go
```

## Gameplay

- The Tic-Tac-Toe board is displayed with numbered cells. Players input the cell number to make their move.
- Player 1 starts with "O", and Player 2 with "X".
- The game checks for a winner after each move and ends when one player achieves three in a row horizontally, vertically, or diagonally, or when all cells are filled with no winner (draw).

## Usage

- Follow the on-screen instructions to play Tic-Tac-Toe.
- Each player takes turns entering a number corresponding to the cell they wish to mark.
- To make a move, enter the number corresponding to the cell where you want to place your sign.
- The game notifies the players when there is a winner or if the game ends in a draw.

<!-- ## Video

- <img src="Video/game-play-tic-tac-toe.mp4" height="400" width="600" alt="Watch Gameplay Video" /> -->

## Contributing

- Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.

## Feedback

- If you have any feedback, please reach out to me at manankoyawala.dev@gmail.com

## Authors

- [@mananKoyawala](https://github.com/mananKoyawala)

## License

[MIT License](LICENSE)
