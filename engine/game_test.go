package engine

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initGame() (*sync.WaitGroup, chan bool, *Game) {
	wg := &sync.WaitGroup{}
	end := make(chan bool)
	mu := &sync.RWMutex{}
	nmap := make(map[string]int)
	kvc := make(chan map[string]int)
	game := NewGame(end, mu, nmap, wg, kvc)
	return wg, end, game
}

func TestGame_RandomPlayer(t *testing.T) {
	wg, end, game := initGame()
	wg.Add(1)
	go game.RandomPlayer(42, 66)
	no := <-end
	wg.Wait()
	wg.Add(1)
	go game.RandomPlayer(42, 42)
	yes := <-end
	wg.Wait()
	assert.Equal(t, no, false)
	assert.Equal(t, yes, true)

}

func TestGame_RememberPlayer(t *testing.T) {
	wg, end, game := initGame()
	wg.Add(1)
	go game.RememberPlayer(42, 33)
	no := <-end
	wg.Wait()
	wg.Add(1)
	go game.RandomPlayer(42, 42)
	yes := <-end
	wg.Wait()
	assert.Equal(t, no, true)
	assert.Equal(t, yes, true)

}

func TestGame_ThoroughPlayer(t *testing.T) {
	wg, end, game := initGame()
	wg.Add(1)
	go game.ThoroughPlayer(42, 42)
	yes := <-end
	wg.Wait()
	//assert.Equal(t, no, true)
	assert.Equal(t, yes, true)

}

func TestGame_CheaterPlayer(t *testing.T) {
	wg, end, game := initGame()
	wg.Add(1)
	go game.CheaterPlayer(42, 42)
	yes := <-end
	wg.Wait()
	//assert.Equal(t, no, true)
	assert.Equal(t, yes, true)

}
