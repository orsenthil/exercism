//
// This is only a SKELETON file for the 'Bank Account' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class ValueError extends Error {
  constructor() {
    super('Bank account error')
  }
}

export class BankAccount {
  currentBalance: number

  constructor() {
    this.currentBalance = NaN;
  }

  open(): void {
    if (!isNaN(this.currentBalance)) {
      throw new ValueError();
    }
    this.currentBalance = 0;
  }

  close(): void {
    if (isNaN(this.currentBalance)) {
      throw new ValueError();
    }
    this.currentBalance = NaN;
  }

  deposit(amount: number): void {
    if (isNaN(this.currentBalance) || amount < 0) {
      throw new ValueError();
    }
    this.currentBalance += amount;
  }

  withdraw(amount: number): void{
    if (isNaN(this.currentBalance) || amount < 0 || amount > this.currentBalance) {
      throw new ValueError();
    }
    this.currentBalance -= amount;
  }

  get balance(): unknown {
    if (isNaN(this.currentBalance)) {
      throw new ValueError();
    }
    return this.currentBalance;
  }
}
