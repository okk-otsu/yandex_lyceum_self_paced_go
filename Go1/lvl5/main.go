package main

import (
	"fmt"
	"sort"
)

type Team struct {
	players []Player
}
type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (p *Player) calculateRating() {
	if p.Misses == 0 {
		p.Rating = float64(p.Goals) + float64(p.Assists)/2
	} else {
		p.Rating = (float64(p.Goals) + float64(p.Assists)/2) / float64(p.Misses)
	}
}
func NewPlayer(name string, goals, misses, assists int) Player {
	p := Player{name, goals, misses, assists, 0.0}
	p.calculateRating()
	return p
}

func goalsSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Goals == players[j].Goals {
			return players[i].Name < players[j].Name
		}
		return players[i].Goals > players[j].Goals
	})
	return players
}

func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Rating == players[j].Rating {
			return players[i].Name < players[j].Name
		}
		return players[i].Rating > players[j].Rating
	})
	return players
}

func gmSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		ri := float64(players[i].Goals) / float64(players[i].Misses)
		rj := float64(players[j].Goals) / float64(players[j].Misses)
		if ri == rj {
			return players[i].Name < players[j].Name
		}
		return ri > rj
	})
	return players
}
func main() {
	players := []Player{
		NewPlayer("Alice", 10, 2, 5),
		NewPlayer("Bob", 15, 3, 2),
		NewPlayer("Charlie", 10, 1, 4),
	}
	sorted := ratingSort(players)
	fmt.Println(sorted)
}
