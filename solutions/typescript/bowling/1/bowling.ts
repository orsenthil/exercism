export class Bowling {
  private round = 0;
  private frames: number[] = [];
  private lastRolls: number[] = [];

  public roll(pins: number): void {
    if (pins < 0) throw new Error("Negative roll is invalid");
    if (pins > 10) throw new Error("Pin count exceeds pins on the lane");
    if (this.round >=  10) throw new Error("Cannot roll after game is over");
    const isStrike = this.lastRolls[0] === 10,
          isSpare = this.lastRolls[0] + this.lastRolls[1] === 10,
          isOpen = this.lastRolls[0] > -1;
    if (isStrike) {
      this.frames[this.round++] = this.lastRolls[0] + this.lastRolls[1] + pins;
      this.lastRolls = [this.lastRolls[1], pins]
    } else if (isSpare) {
      this.frames[this.round++] = this.lastRolls[0] + this.lastRolls[1] + pins;
      this.lastRolls = [-1, pins];
    } else if (isOpen) {
      this.frames[this.round++] = this.lastRolls[0] + this.lastRolls[1];
      this.lastRolls = [-1, pins];
    } else {
      const potentialLastFrame = this.lastRolls[1] + pins;
      if (this.round >= 9 && this.lastRolls[1] > -1 && potentialLastFrame < 10) {
        this.frames[this.round++] = potentialLastFrame
      }
      this.lastRolls = [this.lastRolls[1], pins];
    }

    if (this.lastRolls[0] > 0 && this.lastRolls[0] < 10 && this.lastRolls[0] + this.lastRolls[1] > 10) {
      throw new Error("Pin count exceeds pins on the lane");
    }
  }
  public score(): number {
    if (this.round <= 9) throw new Error("Score cannot be taken until the end of the game");
    return this.frames.reduce((acc, frame) => acc + frame);
  }
}
