export class Clock {

  hour: number = 0
  minute: number = 0

  constructor(hour: number, minute?: number) {
      this.hour = hour
      if (hour >= 24) {
          this.hour = hour % 24
      } else if (hour < 0) {
          while (this.hour < 0) {
              this.hour = 24 + this.hour
          }
      }
      if (minute) {
          if (minute >= 60) {
              this.hour += Math.floor(minute / 60)
              if (this.hour >= 24) {
                  this.hour = this.hour % 24
              }
              this.minute = minute % 60
          } else if (minute < 0) {
              this.hour -= Math.ceil(Math.abs(minute) / 60)
              while (this.hour < 0) {
                  this.hour = 24 + this.hour
              }
              this.minute = 60 - Math.abs(minute) % 60
          } else {
              this.minute = minute
          }
      }
  }

  public toString(): string {
    if (this.hour < 10 && this.minute < 10) {
      return `0${this.hour}:0${this.minute}`
    } else if (this.hour < 10) {
      return `0${this.hour}:${this.minute}`
    } else if (this.minute < 10) {
      return `${this.hour}:0${this.minute}`
    }
    return `${this.hour}:${this.minute}`
  }

  public plus(minutes: number): Clock {
      this.minute += minutes
      if (this.minute >= 60) {
          this.hour += Math.floor(this.minute / 60)
          if (this.hour >= 24) {
              this.hour = this.hour % 24
          }
          this.minute = this.minute % 60
      }
      return this
  }

  public minus(minutes: number): Clock {
      this.minute -= minutes
      if (this.minute < 0) {
          this.hour -= Math.ceil(Math.abs(this.minute) / 60)
          while (this.hour < 0) {
              this.hour = 24 + this.hour
          }
          this.minute = 60 - Math.abs(this.minute) % 60
      }
      return this
  }

  public equals(other: Clock): boolean {
      if (this.hour === other.hour && this.minute === other.minute) {
          return true
      }
      return false
  }
}
