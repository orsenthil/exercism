export class List {
  currentList: number[]

  constructor(...args: number[]) {
    this.currentList = args
  }

  compare(another: List): string {
    if (this.currentList.length === another.currentList.length) {
      if (this.currentList.every((value, index) => value === another.currentList[index])) {
        return 'equal'
      }
    }

    if (this.currentList.length > another.currentList.length) {
      if (this.currentList.join(',').includes(another.currentList.join(','))) {
        return 'superlist'
      }
    }

    if (this.currentList.length < another.currentList.length) {
      if (another.currentList.join(',').includes(this.currentList.join(','))) {
        return 'sublist'
      }
    }

    return 'unequal'
  }
}
