class Node<TElement> {
    value: TElement;
    next: Node<TElement> | null;
    constructor(value: TElement) {
        this.value = value;
        this.next = null;
    }
}

export class LinkedList<TElement> {

  head: Node<TElement> | null;

  constructor() {
      this.head = null;
  }

  public push(element: TElement) {
      const newNode = new Node<TElement>(element);
      if (!this.head) {
          this.head = newNode;
      } else {
          let current = this.head;
          while (current.next) {
              current = current.next;
          }
          current.next = newNode;
      }
  }

  public pop(): TElement {
        let current = this.head;
        let prev = null;
        while (current !== null && current.next !== null) {
            prev = current;
            current = current.next;
        }
        if (prev) {
            prev.next = null;
        } else {
            this.head = null;
        }
        return current.value;
  }

  public shift(): unknown {
    throw new Error('Remove this statement and implement this function')
  }

  public unshift(element: unknown) {
    throw new Error('Remove this statement and implement this function')
  }

  public delete(element: unknown) {
    throw new Error('Remove this statement and implement this function')
  }

  public count(): unknown {
    throw new Error('Remove this statement and implement this function')
  }
}
