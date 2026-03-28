export class Robot {
  robotName: string = '';
  static generatedNames: string[] = [];

  constructor() {
    this.robotName = Robot.generateName();
    Robot.generatedNames.push(this.robotName);
    return this;
  }


  public static generateName(): string {
      let name = '';
      for (let i = 0; i < 2; i++) {
          name += String.fromCharCode(Math.floor(Math.random() * 26) + 65);
      }
      for (let i = 0; i < 3; i++) {
          name += Math.floor(Math.random() * 10);
      }

      if (Robot.generatedNames.includes(name)) {
          return Robot.generateName();
      }

      return name;
  }

  public get name(): string {
    return this.robotName;
  }

  public resetName(): void {
    this.robotName = Robot.generateName();
    Robot.generatedNames.push(this.robotName);
  }

  public static releaseNames(): void {
    Robot.generatedNames = [];
  }
}
