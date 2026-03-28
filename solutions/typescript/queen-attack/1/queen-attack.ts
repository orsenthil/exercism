type Position = readonly [number, number]

type Positions = {
  white: Position
  black: Position
}
export class QueenAttack {
  public readonly black: Position
  public readonly white: Position

  // white: [whiteRow, whiteColumn]
  // black: [blackRow, blackColumn]
  constructor({white = [7, 3], black = [0, 3]}: Partial<Positions> = {}) {
    const onBoard = ([y, x]: Position) => x >= 0 && x <= 7 && y >= 0 && y <= 7;
    if (!onBoard(white) || !onBoard(black)) {
      throw new Error('Queen must be placed on the board');
    }
    if (white[0] === black[0] && white[1] === black[1]) {
      throw new Error('Queens cannot share the same space');
    }

    this.white = white;
    this.black = black;
  }

  toString() {
    let board: string[][] = Array(8).fill(null).map(_ => Array(8).fill('_'));
    const [whiteRow, whiteColumn] = this.white;
    const [blackRow, blackColumn] = this.black;
    board[whiteRow][whiteColumn] = 'W';
    board[blackRow][blackColumn] = 'B';

    return board.map(row => row.join(' ')).join('\n');
  }

  get canAttack() {
    const [whiteRow, whiteColumn] = this.white;
    const [blackRow, blackColumn] = this.black;

    return (
      whiteRow === blackRow ||
      whiteColumn === blackColumn ||
      Math.abs(whiteRow - blackRow) === Math.abs(whiteColumn - blackColumn) // diagonal
    );
  }
}