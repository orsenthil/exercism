class Node {
  constructor(
    public data: number,
    public left: Node | null = null,
    public right: Node | null = null) { }

}
export class BinarySearchTree {
  private root: Node

  constructor(data: number) {
    this.root = new Node(data)
  }

  public get data(): number {
    return this.root.data
  }

  public get right(): Node | null {
    return this.root.right
  }

  public get left(): Node | null {
    return this.root.left
  }

  public insert(item: number, root = this.root): void {
    if (item <= root.data) {
      if (root.left) {
        this.insert(item, root.left)
      } else {
        root.left = new Node(item)
      }
    } else if (item > root.data) {
      if (root.right) {
        this.insert(item, root.right)
      } else {
        root.right = new Node(item)
      }
    }
  }

  public each(callback: (data: number) => void): void {
    this.inOrderTraversal(this.root).forEach(callback)
  }

  private inOrderTraversal(root: Node | null = this.root): number[] {
    if (!root) {
      return []
    }

    return [
      ...this.inOrderTraversal(root.left),
      root.data,
      ...this.inOrderTraversal(root.right),
    ]
  }
}
