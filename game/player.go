package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/jcheng31/diceroller/dice"
)

// Player structure holds relevant data for a player to keep track of during the dice game.
//
// Members:
//
// * `roundScores`: keeps the total for each round
//
// * `currentRoundDice`: shows the current dice rolls for this round (Rolls that are kept are not rerolled)
//
// * `keptDiceRoll`: marks which die of our current round are being kept
//
// * `playerName`: the name to be displayed for the player
type Player struct {
	roundScores      [AmountOfRounds]int
	currentRoundDice [AmountOfDice]int
	keptDiceRoll     [AmountOfDice]bool
	playerName       string
}

//Has a player take their turn, rolls their dice then allows them to choose which they'd like to keep
func (player *Player) takeTurn(roundIndex int) {

	//loop could be removed to allow for each player to take a turn in round, instead of each player exhausting their rolls
	for !player.hasFinishedTurn() {
		fmt.Printf("\n%s's turn", player.playerName)

		//roll current round dice left
		player.populateRoundDice()

		//display to player kept dice & dice pool to choose from
		player.displayCurrentRoundDice()

		//allow the player to keep 1 - n amount of dice, where n is remaining pool to choose from
		player.chooseKeptDice()
	}

	//when all have been chosen, tally total and store it in roundScores
	player.tallyRoundScore(roundIndex)
	player.clearRound()
}

//Asks the player how many dice they'd like to keep out of their remaining pool, then walks them through picking each one to keep
func (player *Player) chooseKeptDice() {
	if player.hasFinishedTurn() {
		return
	}

	numberOfUnkeptDiceRolls := player.getUnkeptDiceRolls()

	if numberOfUnkeptDiceRolls == 1 {
		fmt.Println("Marking last die as kept.")
		player.markAllDiceAsKept()
		return
	}

	fmt.Printf("How many dice rolls would you like to keep? You have %d unkept rolls.\n", numberOfUnkeptDiceRolls)

	choice := ""
	if SimulatePlay {
		choice = "5"
	} else {
		choice = getUserChoice()
	}

	if userChoice, err := strconv.Atoi(choice); err == nil && isValidIndexChoice(userChoice, 1, AmountOfDice) {

		//if the player chooses to keep all the remaining dice.. we don't want to keep asking them to enter the index, just mark them all
		if userChoice == numberOfUnkeptDiceRolls {
			fmt.Println("Marking remaining dice as kept.")
			player.markAllDiceAsKept()

		} else {
			player.exhaustUntilAllChosenDiceKept(userChoice)
		}
	}
}

//Continously ask the player which die they would like to keep until their choosen amount of dice to keep has been met
func (player *Player) exhaustUntilAllChosenDiceKept(userChoice int) {
	for i := 0; i < userChoice; i++ {
		player.keptDiceRoll[player.getValidatedChoiceIndex("Please enter the index of the die roll to keep.")-1] = true
	}
}

//Marks all dice as kept
func (player *Player) markAllDiceAsKept() {
	for i := range player.keptDiceRoll {
		player.keptDiceRoll[i] = true
	}
}

//Returns the amount of dice that have not been marked as kept
func (player *Player) getUnkeptDiceRolls() int {
	tally := 0

	for i := range player.keptDiceRoll {
		if !player.keptDiceRoll[i] {
			tally++
		}
	}

	return tally
}

//Returns whether the given die is kept or not
func (player *Player) isDieKept(index int) bool {
	return player.keptDiceRoll[index]
}

//Get the users choice from console
func getUserChoice() string {
	byteSizeStr := strconv.Itoa(AmountOfDice)
	byteSize := len(byteSizeStr)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	userInput := string([]byte(input))
	choice := string([]byte(input)[byteSize-1])

	difference := (len(userInput) - len(choice)) - 1
	if difference > byteSize {
		fmt.Printf("You've entered: %sAccepting: %s, as the answer.\n", userInput, choice)
	}

	return choice
}

//Returns the validated userChoice, ensuring that it is in the right range
func (player *Player) getValidatedChoiceIndex(printText string) int {
	fmt.Println(printText)

	choice := getUserChoice()

	if num, err := strconv.Atoi(choice); err == nil && isValidIndexChoice(num, 1, AmountOfDice) && !player.isDieKept(num-1) {
		return num
	}

	return player.getValidatedChoiceIndex("Please enter a valid index.")
}

//Utility function that ensures a given number is within a bounds, inclusive
func isValidIndexChoice(choice int, lowBound int, highBound int) bool {
	if choice <= highBound && choice >= lowBound {
		return true
	}

	return false
}

//Display the player's round score
func (player *Player) displayRoundScore(roundIndex int) {
	fmt.Printf("\n%s's score for round %d was %d \n", player.playerName, roundIndex, player.roundScores[roundIndex-1])
}

//Tally the player's round score
func (player *Player) tallyRoundScore(roundIndex int) {
	accumulator := 0
	for i := range player.currentRoundDice {
		if player.currentRoundDice[i] != ZeroNumber {
			accumulator = accumulator + player.currentRoundDice[i]
		}
	}

	player.roundScores[roundIndex-1] = accumulator
}

//Reset the round back to default, no leftover data for the next round
func (player *Player) clearRound() {
	for i := range player.keptDiceRoll {
		player.keptDiceRoll[i] = false
	}

	for i := range player.currentRoundDice {
		player.currentRoundDice[i] = 0
	}
}

//Displays the current kept dice, and the current pool of dice the player has to choose from
func (player *Player) displayCurrentRoundDice() {
	kept := "\nKept: "
	choicePool := "Choice Pool: "

	for i := range player.currentRoundDice {
		if player.keptDiceRoll[i] {
			kept = kept + strconv.Itoa(player.currentRoundDice[i]) + " "
			choicePool = choicePool + "X "
		} else {
			choicePool = choicePool + strconv.Itoa(player.currentRoundDice[i]) + " "
		}
	}

	fmt.Printf("%s\n%s\n", kept, choicePool)
}

//Checks to see if the player has finished their turn, but seeing if they have kept all their dice or not
func (player *Player) hasFinishedTurn() bool {
	for i := range player.keptDiceRoll {
		if !player.keptDiceRoll[i] {
			return false
		}
	}

	return true
}

//Rolls the dice for the current turn
func (player *Player) populateRoundDice() {
	for i := range player.currentRoundDice {
		if !player.keptDiceRoll[i] {
			player.currentRoundDice[i] = getDieRoll(MaxDieFace)
		}
	}
}

//Adds together this players round scores, returning the total
func (player *Player) calcuateTotalGameScore() int {
	totalScore := 0

	for i := range player.roundScores {
		totalScore += player.roundScores[i]
	}

	return totalScore
}

//Utility function, used for rolling a dice
func getDieRoll(diceFace int) int {

	d6 := dice.Regular(RandomRoller, diceFace)
	result := d6.RollN(1)
	return result.Rolls[0]
}
