export default class CircularBuffer<T> {
  private readonly size: Number
  private store: T[] = []
  
  constructor(initial: Number) {
    this.size = initial
  }

  write(value: T): void {
    if (CircularBuffer.full(this.store, this.size)) {
      throw new BufferFullError();
    }
    this.store.push(value);
  }

  read(): T | undefined{
    if (CircularBuffer.empty(this.store)) {
      throw new BufferEmptyError();
    }
    return this.store.shift();
  }

  forceWrite(value: T): void {
    if (CircularBuffer.full(this.store, this.size)) {
      this.read();
    }
    this.write(value);
  }

  clear(): void {
    this.store = [];
  }

  private static empty<T>(store: T[]): boolean {
    return store.length === 0;
  }

  private static full<T>(store: T[], size: Number): boolean {
    return store.length === size;
  }

}

export class BufferFullError extends Error {
  constructor() {
    super()
  }
}

export class BufferEmptyError extends Error {
  constructor() {
    super()
  }
}
