package prompt

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/miko-code/gw/model"
	"go.uber.org/zap"
)

type Prompt struct{}

func NewPrompt() *Prompt {
	return &Prompt{}
}

//TO DO VALIDATE NUMBERS 2-8
func getNumOfPlayers() float64 {
	validate := func(input string) error {
		n, err := strconv.ParseFloat(input, 64)
		if err != nil || n > 8 {
			return errors.New("Invalid number ")
		}

		return nil
	}
	t := promptui.Prompt{
		Label:    "Number of players",
		Validate: validate,
	}
	n, _ := t.Run()
	numberOfPlayers, _ := strconv.ParseFloat(n, 64)
	return numberOfPlayers

}

func getname() string {
	prompt := promptui.Prompt{
		Label: "Player Name",
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err.Error())
	}
	return result
}

func getType() string {
	selectType := promptui.Select{
		Label: "Select Player Character",
		Items: []string{"random_player", "remember_player", "thorough_player", "cheater_player", "thorough_cheater_player"},
	}
	_, selectResult, err := selectType.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

	}
	return selectResult
}

func (p *Prompt) GeneratePlayersInfo() *[]model.Player {
	n := getNumOfPlayers()

	players := make([]model.Player, int(n))

	for i := 0; i < int(n); i++ {
		name := getname()
		t := getType()
		player := model.NewPlayer(name, t, i)
		players = append(players, player)
		zap.S().Info(fmt.Sprintf("player info %d,%s,%s", i, name, t))
	}
	return &players
}
