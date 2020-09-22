package model

func NewPlayer(name string, character string, id int) Player {
	p := Player{
		Name:      name,
		Character: character,
		Id:        id,
	}
	return p
}

type Player struct {
	Name      string
	Character string
	Id        int
}
