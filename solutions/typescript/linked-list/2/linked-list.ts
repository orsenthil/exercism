class Node<TElement> {
    value: TElement;
    next: Node<TElement> | null;
    constructor(value: TElement) {
        this.value = value;
        this.next = null;
    }
}

export class LinkedList<TElement> {

    list: Node<TElement>[];

    public constructor() {
        this.list = [];
    }

  public push(element: TElement) {
      const newElement: Node<TElement> = new Node(element);
      if (this.list.length === 0) {
          this.list.push(newElement);
      } else {
          this.list[this.list.length - 1].next = newElement;
          this.list.push(newElement);
      }
  }

  public pop(): TElement | null {
      const lastElement = this.list.pop();
      if (this.list.length > 0) {
          this.list[this.list.length - 1].next = null;
      }

      return lastElement ? lastElement.value : null;
  }

  public shift(): TElement | null {
      const firstElement = this.list.shift();
      return firstElement ? firstElement.value : null;
  }

  public unshift(element: TElement) {
      const newElement: Node<TElement> = new Node(element);
      this.list.unshift(newElement);
      if (this.list.length > 1) {
          newElement.next = this.list[1];
      }
  }

  public delete(element: TElement) {
      const index = this.list.findIndex((node) => node.value === element);
      if (index > -1) {
          if (index - 1 >= 0) {
              this.list[index - 1].next = this.list[index + 1];
          }
          this.list.splice(index, 1);
      }
  }

  public count(): unknown {
      return this.list.length;
  }
}
