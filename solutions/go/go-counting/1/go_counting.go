package gocounting

import "errors"

type AllTerritories struct {
	Black [][2]int
	White [][2]int
	None  [][2]int
}

type Game struct {
	board []string
}

func NewGame(board []string) *Game {
	return &Game{board: board}
}

func (g *Game) Territory(x, y int) (string, [][2]int, error) {
	if x < 0 || x >= len(g.board[0]) || y < 0 || y >= len(g.board) {
		return "", nil, errors.New("invalid coordinate")
	}
	if g.board[y][x] == 'B' || g.board[y][x] == 'W' {
		return "NONE", [][2]int{}, nil
	}
	// flood fill
	visited := map[[2]int]bool{}
	queue := [][2]int{{x, y}}
	territory := [][2]int{}
	seenBlack := false
	seenWhite := false

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		visited[curr] = true
		for _, neighbor := range getNeighbors(g.board, curr) {
			if !visited[neighbor] {
				if g.board[neighbor[1]][neighbor[0]] == ' ' {
					queue = append(queue, neighbor)
					visited[neighbor] = true
				}
				if g.board[neighbor[1]][neighbor[0]] == 'B' {
					seenBlack = true
				}
				if g.board[neighbor[1]][neighbor[0]] == 'W' {
					seenWhite = true
				}
			}
		}
	}
	for pos := range visited {
		territory = append(territory, pos)
	}
	if seenBlack && seenWhite {
		return "NONE", territory, nil
	}
	if seenBlack {
		return "BLACK", territory, nil
	}
	if seenWhite {
		return "WHITE", territory, nil
	}
	return "NONE", territory, nil
}

func getNeighbors(board []string, pos [2]int) [][2]int {
	x, y := pos[0], pos[1]
	neighbors := [][2]int{}
	for _, neighbor := range [][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}} {
		if neighbor[0] >= 0 && neighbor[0] < len(board[0]) && neighbor[1] >= 0 && neighbor[1] < len(board) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func (g *Game) Territories() AllTerritories {
	territories := AllTerritories{
		Black: [][2]int{},
		White: [][2]int{},
		None:  [][2]int{},
	}
	for y := range g.board {
		for x := range g.board[y] {
			if g.board[y][x] == ' ' {
				owner, _, err := g.Territory(x, y)
				if err != nil {
					continue
				}
				switch owner {
				case "BLACK":
					territories.Black = append(territories.Black, [2]int{x, y})
				case "WHITE":
					territories.White = append(territories.White, [2]int{x, y})
				case "NONE":
					territories.None = append(territories.None, [2]int{x, y})
				}
			}
		}
	}
	return territories
}
