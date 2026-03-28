// Package connect provides functionality for determining winners of Hex/Poligon games.
package connect
var adjacentHexes = [6][2]int{{0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}}
type node [2]int
type route struct {
	isDeadEnd bool
	nodes     []node
}
// ResultOf determines the winner of the given board.
func ResultOf(input []string) (result string, err error) {
	players := [2]rune{'O', 'X'}
	board := make([][]rune, 0, len(input))
	for i := 0; i < len(input); i++ {
		board = append(board, []rune(input[i]))
	}
	for i := 0; i < 2; i++ {
		startNodes := getStartNodes(players[i], board)
		if isWin(players[i], board, startNodes) {
			return string(players[i]), nil
		}
	}
	return "", nil
}
func getStartNodes(player rune, board [][]rune) []node {
	var result []node
	if player == 'O' {
		for x := 0; x < len(board[0]); x++ {
			if board[0][x] == player {
				result = append(result, [2]int{x, 0})
			}
		}
	} else if player == 'X' {
		for y := 0; y < len(board); y++ {
			if board[y][0] == player {
				result = append(result, [2]int{0, y})
			}
		}
	}
	return result
}
func getNextNodes(n node, player rune, board [][]rune, rte route) []node {
	var result []node
	for _, v := range adjacentHexes {
		x, y := n[0]+v[0], n[1]+v[1]
		if 0 <= x && x < len(board[0]) && 0 <= y && y < len(board) && board[y][x] == player && !isCircularRoute(x, y, rte) {
			result = append(result, [2]int{x, y})
		}
	}
	return result
}
func isWin(player rune, board [][]rune, startNodes []node) bool {
	for _, n := range startNodes {
		routes := []route{route{isDeadEnd: false, nodes: []node{n}}}
		isConnected := isAnyRouteConnected(player, routes, board)
		isDeadEnd := areAllRoutesDeadEnd(routes)
		var routesToAdd []route
		for isConnected == false && isDeadEnd == false {
			routesToAdd = make([]route, 0)
			for i := range routes {
				if routes[i].isDeadEnd == false {
					nextNodes := getNextNodes(routes[i].nodes[len(routes[i].nodes)-1], player, board, routes[i])
					if len(nextNodes) == 0 {
						routes[i].isDeadEnd = true
					} else {
						for j := len(nextNodes) - 1; j >= 0; j-- {
							if j == 0 {
								routes[i].nodes = append(routes[i].nodes, nextNodes[j])
							} else {
								tmpNodes := make([]node, len(routes[i].nodes))
								copy(tmpNodes, routes[i].nodes)
								tmpNodes = append(tmpNodes, nextNodes[j])
								routesToAdd = append(routesToAdd, route{isDeadEnd: false, nodes: tmpNodes})
							}
						}
					}
				}
			}
			routes = append(routes, routesToAdd...)
			isConnected = isAnyRouteConnected(player, routes, board)
			isDeadEnd = areAllRoutesDeadEnd(routes)
		}
		if isConnected {
			return true
		}
	}
	return false
}
func isAnyRouteConnected(player rune, routes []route, board [][]rune) bool {
	for _, rte := range routes {
		if player == 'O' && rte.nodes[len(rte.nodes)-1][1] == len(board)-1 {
			return true
		} else if player == 'X' && rte.nodes[len(rte.nodes)-1][0] == len(board[0])-1 {
			return true
		}
	}
	return false
}
func areAllRoutesDeadEnd(routes []route) bool {
	for _, rte := range routes {
		if rte.isDeadEnd == false {
			return false
		}
	}
	return true
}
func isCircularRoute(x, y int, rte route) bool {
	for _, v := range rte.nodes {
		if x == v[0] && y == v[1] {
			return true
		}
	}
	return false
}
