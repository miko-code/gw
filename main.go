package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/miko-code/gw/conf"
	"github.com/miko-code/gw/engine"
	"github.com/miko-code/gw/model"
	"github.com/miko-code/gw/prompt"
	"go.uber.org/zap"
)

func main() {
	logger := initLogger()
	logger.Info("Game start")
	config := conf.GetConf()
	logger.Info("iteration ", config.Iteration)
	players, wg, end, mu, nmap, kvc := initGame()
	game := engine.NewGame(end, mu, nmap, wg, kvc)

	n := game.GenrateNumber()
	fmt.Println("Sack wight is  ", n)
	logger.Info("Sack wight is  ", n)

	wg.Add(config.Iteration)

	for i := 1; i <= config.Iteration; i++ {
		game.TryGuess(players, n)
		go func() {
			select {
			case gameStatus := <-end:
				if gameStatus {
					os.Exit(0)
				} else {
					logger.Info("Game over no winner")

				}

				return
			}
		}()
	}
	wg.Wait()

}

func initLogger() *zap.SugaredLogger {
	l := conf.NewLogger()
	zlog, err := l.GetLogger()
	if err != nil {
		panic(err.Error())
	}
	zap.ReplaceGlobals(zlog)
	defer zlog.Sync()
	logger := zlog.Sugar()
	return logger
}

func initGame() (*[]model.Player, *sync.WaitGroup, chan bool, *sync.RWMutex, map[string]int, chan map[string]int) {
	prm := prompt.NewPrompt()
	players := prm.GeneratePlayersInfo()
	wg := &sync.WaitGroup{}
	end := make(chan bool)
	mu := &sync.RWMutex{}
	nmap := make(map[string]int)
	kvc := make(chan map[string]int)
	rand.Seed(time.Now().UnixNano())
	return players, wg, end, mu, nmap, kvc
}
