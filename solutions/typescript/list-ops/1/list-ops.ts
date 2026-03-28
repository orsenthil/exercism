interface Link {
  value: unknown;
  next: Link | null;
}

export class List {
  protected head: Link | null = null;

  public static create(...values: unknown[]): List {
    const list = new List();
    for (const value of values) {
      list.append(value);
    }
    return list;
  }

  public append(value: unknown): void {
    if (this.head == null) {
      this.head = { value, next: null };
    } else {
      let current: Link | null = this.head;
      while (current.next != null) {
        current = current.next;
      }
      current.next = { value, next: null };
    }
  }

  public concatenate(other: List): void {
    let current: Link | null = this.head;
    if (current == null) {
      this.head = other.head;
    } else {
      while (current.next != null) {
        current = current.next;
      }
      current.next = other.head;
    }
  }

  public filter(predicate: (value: unknown) => boolean): List {
    const list = new List();
    let current: Link | null = this.head;
    while (current != null) {
      if (predicate(current.value)) {
        list.append(current.value);
      }
      current = current.next;
    }
    return list;
  }

  public length(): number {
    let current: Link | null = this.head;
    let length: number = 0;
    while (current) {
      length ++;
      current = current.next;
    }
    return length;
  }

  public map(func: (value: unknown) => unknown): List {
    const list = new List();
    let current: Link | null = this.head;
    while (current != null) {
      list.append(func(current.value));
      current = current.next;
    }
    return list;
  }

  public foldl(func: (acc: unknown, value: unknown) => unknown, zero: unknown): unknown {
    let current: Link | null = this.head;
    let result: unknown = zero;
    while (current != null) {
      result = func(result, current.value);
      current = current.next;
    }
    return result;
  }

  public foldr(func: (acc: unknown, value: unknown) => unknown, zero: unknown): unknown {
    const reversed = this.reverse();
    return reversed.foldl(func, zero);
  }

  public reverse(): List {
    let current: Link | null = this.head;
    const list = new List();
    while (current != null) {
      const oldHead = list.head;
      list.head = { value: current.value, next: oldHead };
      current = current.next;
    }
    return list;
  }
}
