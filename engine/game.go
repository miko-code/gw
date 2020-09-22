package engine

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/miko-code/gw/model"
	"go.uber.org/zap"
)

type Game struct {
	End  chan bool
	Mu   *sync.RWMutex
	Nmap map[string]int
	Wg   *sync.WaitGroup
	Kvc  chan map[string]int
}

//NewGame constractor
func NewGame(
	end chan bool, mu *sync.RWMutex, nmap map[string]int, wg *sync.WaitGroup, kvc chan map[string]int) *Game {
	return &Game{
		End:  end,
		Mu:   mu,
		Nmap: nmap,
		Wg:   wg,
		Kvc:  kvc,
	}
}

//RandomPlayer try to guess  the number
func (g *Game) RandomPlayer(n int, gn int) {
	defer g.Wg.Done()
	g.setusedNum(fmt.Sprintf("%s_%d", "All", gn), gn)
	zap.S().Info(fmt.Sprintf("player RandomPlayer try  %d,", gn))

	if gn == n {
		zap.S().Info(fmt.Sprintf("player RandomPlayer is the Winner!!!"))
		g.End <- true
		return
	}
	t := g.getAbsNum(gn, n)
	time.Sleep(time.Duration(t) * time.Millisecond)
	g.End <- false

}

//RememberPlayer try to guess  the number
func (g *Game) RememberPlayer(n int, gn int) {
	defer g.Wg.Done()
	for {
		_, ok := g.getusedNum(fmt.Sprintf("%s_%d", "Rem", gn))
		if !ok {
			break
		}
		gn = g.GenrateNumber()
	}
	zap.S().Info(fmt.Sprintf("player RememberPlayer try  %d,", gn))

	g.setusedNum(fmt.Sprintf("%s_%d", "Rem", gn), gn)
	if gn == n {

		zap.S().Info(fmt.Sprintf("player RememberPlayer is the Winner!!!"))
		g.End <- true
		return
	}
	t := g.getAbsNum(gn, n)
	time.Sleep(time.Duration(t) * time.Millisecond)
	g.End <- false
}

//ThoroughPlayer try to guess  the number
func (g *Game) ThoroughPlayer(i int, n int) {
	defer g.Wg.Done()
	zap.S().Info(fmt.Sprintf("player ThoroughPlayer try  %d,", i))

	g.setusedNum(fmt.Sprintf("%s_%d", "All", i), i)
	if i == n {
		zap.S().Info(fmt.Sprintf("player ThoroughPlayer is the Winner!!!"))
		g.End <- true
		return
	}
	t := g.getAbsNum(i, n)
	time.Sleep(time.Duration(t) * time.Millisecond)
	g.End <- false

}

//CheaterPlayer try to guess  the number
func (g *Game) CheaterPlayer(n int, gn int) {

	for {
		_, ok := g.getusedNum(fmt.Sprintf("%s_%d", "All", gn))
		if !ok {
			break
		}
		gn = g.GenrateNumber()
	}
	zap.S().Info(fmt.Sprintf("player CheaterPlayer try  %d,", gn))

	g.setusedNum(fmt.Sprintf("%s_%d", "All", gn), gn)
	if gn == n {
		zap.S().Info(fmt.Sprintf("player CheaterPlayer is the Winner!!!"))
		g.End <- true

		return
	}
	t := g.getAbsNum(gn, n)
	time.Sleep(time.Duration(t) * time.Millisecond)
	g.End <- false
}

//ThoroughCheaterPlayer try to guess  the number
func (g *Game) ThoroughCheaterPlayer(i int, n int) {
	zap.S().Info(fmt.Sprintf("player ThoroughCheaterPlayer try  %d,", i))
	g.CheaterPlayer(n, i)
}

//TryGuess wrapper iterate over players and execute goutines with new numbers
func (g *Game) TryGuess(players *[]model.Player, n int) {

	for _, player := range *players {

		switch player.Character {
		case "random_player":
			gn := g.GenrateNumber()
			go g.RandomPlayer(n, gn)
		case "remember_player":
			gn := g.GenrateNumber()
			go g.RememberPlayer(n, gn)
		case "thorough_player":
			for i := 40; i <= 140; i++ {
				go func() {
					g.ThoroughPlayer(n, i)
				}()
			}
		case "cheater_player":
			gn := g.GenrateNumber()
			go g.CheaterPlayer(n, gn)
		case "thorough_cheater_player":
			for i := 40; i <= 140; i++ {
				go func() {
					g.ThoroughCheaterPlayer(n, i)
				}()
			}

		}

	}

}

//GenrateNumber create new number
func (g *Game) GenrateNumber() int {

	return rand.Intn(140-40) + 40
}

func (g *Game) setusedNum(k string, num int) {
	g.Mu.Lock()
	s := fmt.Sprintf("%s_%d", k, num)
	g.Nmap[s] = num
	g.Mu.Unlock()

}
func (g *Game) getusedNum(k string) (int, bool) {
	g.Mu.RLock()
	n, ok := g.Nmap[k]
	g.Mu.RUnlock()
	return n, ok
}
func (g *Game) getAbsNum(gn int, n int) int {
	x := gn - n
	if x < 0 {
		return -x
	}
	return x
}

//TODO
func (g *Game) clossetPlayer(n int, gn int, p string) {

}
