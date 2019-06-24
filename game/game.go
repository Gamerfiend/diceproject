package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jcheng31/diceroller/roller"
)

//Game Rules
const (
	AmountOfRounds         = 4
	AmountOfPlayers        = 4
	AmountOfDice           = 5
	MaxDieFace             = 6
	ZeroNumber             = 4
	MaxRoundScore          = AmountOfDice * MaxDieFace
	MaxPossibleScore       = MaxRoundScore * AmountOfRounds
	SimulatePlay           = true
	SimulatePlayRandomSeed = 42
)

//RandomRoller Roller, will use a seed if on simulate mode
var RandomRoller = roller.WithRandomSource(rand.NewSource(time.Now().UnixNano()))

// Game structure holds all data relevant to game
//
// Members:
//
// * `players`: an array of pointers to the players
//
// * `playerHasStarted`: an array of boolean values to keep track of if a player has gone first yet
//
// * `roundWinner`: an int array which marks which player won each round
//
// * `roundIndex`: marks which round the game is currently on.. NOTE: this starts at 1 NOT 0
//
// * `gameWinner`: marks the index of the player from players who won the game
type Game struct {
	players          [AmountOfPlayers]*Player
	playerHasStarted [AmountOfPlayers]bool
	roundWinner      [AmountOfRounds]int
	roundIndex       int
	gameWinner       int
}

//Play Begins the game by first getting the player names, then runs through the rounds until it ends with displaying the winner
func (game *Game) Play() {
	fmt.Println("**(>^.^)>**Welcome to Tyler's Game of Dice**<(^.^<)**\n*****************************************************")

	if SimulatePlay {
		fmt.Println("***********************\nGame in Simulation Mode\n***********************")
		game.setPlayerNames()
	} else {
		game.retrieveAndSetPlayerNames()
	}

	for i := 0; i < AmountOfRounds; i++ {
		game.playRound()
	}
	game.calculateGameWinner()
	game.displayGameWinner()
}

//NewGame Constructor function for Game struct, initalizes the array of our players
func (game *Game) NewGame() {

	for i := range game.players {
		game.players[i] = new(Player)
	}
}

//Sets the player names for the game
func (game *Game) retrieveAndSetPlayerNames() {
	for i := range game.players {
		fmt.Printf("Enter name for player %d:\n", i+1)
		var name string
		fmt.Scanln(&name)
		game.players[i].playerName = name
		fmt.Println("")
	}
}

func (game *Game) setPlayerNames() {
	testPlayerNames := [4]string{"Tyler", "David", "Joe", "Tom"}

	for index, name := range testPlayerNames {
		game.players[index].playerName = name
	}
}

//Calculates the game winner
func (game *Game) calculateGameWinner() {
	winningPlayer := 0
	winningTotal := MaxPossibleScore

	for i := range game.players {
		playerTotal := game.players[i].calcuateTotalGameScore()

		if playerTotal < winningTotal {
			winningTotal = playerTotal
			winningPlayer = i
		}
	}

	game.gameWinner = winningPlayer
}

//Displays the game winner
func (game *Game) displayGameWinner() {
	winnerName := game.players[game.gameWinner].playerName
	winnerScore := game.players[game.gameWinner].calcuateTotalGameScore()
	fmt.Printf("\n%s has won the game with a score of %d!!", winnerName, winnerScore)
}

//Determine who won the round
func (game *Game) calculateRoundWinner() {
	lowestScoreIndex := 0 //default to first player.. incase they all have the same score
	lowestScore := MaxRoundScore

	for i := range game.players {
		if game.players[i].roundScores[game.roundIndex] < lowestScore {
			lowestScoreIndex = i
			lowestScore = game.players[i].roundScores[game.roundIndex]
		}
	}

	game.roundWinner[game.roundIndex] = lowestScoreIndex
}

//Display who won the round
func (game *Game) displayRoundWinner() {
	winner := game.players[game.roundWinner[game.roundIndex]]
	fmt.Printf("\nThe winner for round %d, with a score of %d, is %s\n", game.roundIndex+1, winner.roundScores[game.roundIndex], winner.playerName)
}

//Determines and returns the player order for the round of play.
//The function will first randomly find a player who hasn't gone first.
//Then it will set them as first, and randomly populate the rest of the turn
//order until array is full
func (game *Game) getRoundPlayerOrder() [AmountOfPlayers]int {
	playerOrder := [AmountOfPlayers]int{}

	for i := range playerOrder {
		playerOrder[i] = -1
	}

	randomPlayer := getDieRoll(AmountOfPlayers) - 1

	for game.hasPlayerGoneFirst(randomPlayer) {
		randomPlayer = getDieRoll(AmountOfPlayers) - 1
	}

	playerOrder[0] = randomPlayer
	game.playerHasStarted[randomPlayer] = true

	for i := 1; i < AmountOfPlayers; i++ {
		randomPlayer = getDieRoll(AmountOfPlayers) - 1
		for contains(playerOrder, randomPlayer) {
			randomPlayer = getDieRoll(AmountOfPlayers) - 1
		}

		playerOrder[i] = randomPlayer
	}

	return playerOrder
}

//Utility function to determine whether or not the player array contains a player
func contains(playerArray [AmountOfPlayers]int, player int) bool {
	for i := 0; i < AmountOfPlayers; i++ {
		if playerArray[i] == player {
			return true
		}
	}

	return false
}

//Checks to see if given player has gone first
func (game *Game) hasPlayerGoneFirst(index int) bool {
	return game.playerHasStarted[index]
}

//Plays out a round by having each player take their turn, displays their scores, determines the winner and displays it
func (game *Game) playRound() {
	fmt.Printf("\n********Round %d********", game.roundIndex+1)

	playerOrder := game.getRoundPlayerOrder()

	//players take turns
	for i := range playerOrder {
		game.players[playerOrder[i]].takeTurn(game.roundIndex + 1)
	}

	for i := range game.players {
		game.players[playerOrder[i]].displayRoundScore(game.roundIndex + 1)
	}

	//determine winner, and mark it
	game.calculateRoundWinner()
	game.displayRoundWinner()
	game.roundIndex++
}
