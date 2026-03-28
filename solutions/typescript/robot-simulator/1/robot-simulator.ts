export class InvalidInputError extends Error {
  constructor(message: string) {
    super()
    this.message = message || 'Invalid Input'
  }
}

type Direction = 'north' | 'east' | 'south' | 'west'
type Coordinates = [number, number]

interface ManualEntry {
  left: Direction
  right: Direction
  advance: number[]
}

interface Manual {
  [key: string]: ManualEntry
}

const robotManual: Manual = {
  north: {
    left: 'west',
    right: 'east',
    advance: [0, 1],
  },
  east: {
    left: 'north',
    right: 'south',
    advance: [1, 0],
  },
  south: {
    left: 'east',
    right: 'west',
    advance: [0, -1],
  },
  west: {
    left: 'south',
    right: 'north',
    advance: [-1, 0],
  },
}


export class Robot {
  direction: Direction
  coord: Coordinates

  constructor() {
    this.direction = 'north'
    this.coord = [0, 0]
  }

  get bearing(): Direction {
    return this.direction
  }

  get coordinates(): Coordinates {
    return this.coord
  }

  place(placement: { x: number; y: number; direction: string }) {
    if (!['north', 'east', 'south', 'west'].includes(placement.direction)) {
      throw new InvalidInputError('Invalid direction')
    }
    this.direction = placement.direction as Direction
    this.coord = [placement.x, placement.y]
  }

  evaluate(instructions: string) {
    var _this = this;
    [...instructions].forEach(function (value: string) {
      _this.evaluateSingle(value)
    });
  }

  evaluateSingle(instruction: string) {
    switch(instruction) {
      case 'A': this.move(); break;
      case 'L': this.direction = robotManual[this.direction].left; break;
      case 'R': this.direction = robotManual[this.direction].right; break; 
    }
  }

  move() {
    const movement = robotManual[this.direction].advance
    this.coord[0] += movement[0]
    this.coord[1] += movement[1]
  }

}
