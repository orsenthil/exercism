export class Gigasecond {
  inputDate: Date;
  constructor(inputDate: Date) {
    this.inputDate = inputDate;
  }
  public date(): Date {
    return new Date(this.inputDate.getTime() + 1e12);
  }
}
