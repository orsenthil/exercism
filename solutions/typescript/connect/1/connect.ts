const ADJACENTS = [[1, 0], [-1, 0], [0, 1], [0, -1], [1, -1], [-1, 1]]
const X = 'X'
const O = 'O'

export class Board {
  private board: string[][]

  constructor(board: string[]) {
    this.board = board.map(row => row.trim().split(' '))
  }

  public winner(): unknown {
    return this.board.some((_, i) => this.isConnected(X, i, 0)) ? X:
      this.board.some((_, i) => this.isConnected(O, 0, i)) ? O: ''
  }

  private isConnected(player: string, x: number, y: number): boolean {
    switch( true ) {
      case this.board[x][y] !== player: return false
      case player === O && x === this.board.length - 1: return true
      case player === X && y === this.board[0].length - 1: return true
    }
    this.board[x][y] = ''
    return ADJACENTS.some(([dx, dy]) => this.board[x + dx] && this.board[x + dx][y + dy] && this.isConnected(player, x + dx, y + dy))
  }
}
