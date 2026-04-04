package zipper

import "errors"

type Node struct {
	value int
	left  *Node
	right *Node
}

type breadcrumb struct {
	parent *Node
	sibling *Node
	fromLeft bool
}

type Zipper struct {
	current *Node
	trail []breadcrumb
	modified bool
}

func NewZipper(tree *Node) Zipper {
	return Zipper{
		current: tree,
		trail:   []breadcrumb{},
		modified: false,
	}
}

func (z Zipper) Value() int {
	return z.current.value
}

func (z Zipper) ToTree() *Node {
	for len(z.trail) > 0 {
		z, _ = z.Up()
	}
	return z.current
}

func (z Zipper) Left() (Zipper, error) {
	if z.current.left == nil {
		return Zipper{}, errors.New("no left child")
	}
	crumb := breadcrumb{
		parent: z.current,
		sibling: z.current.right,
		fromLeft: true,
	}
	return Zipper{
		current: z.current.left,
		trail: append(z.trail, crumb),
	}, nil
}

func (z Zipper) Right() (Zipper, error) {
	if z.current.right == nil {
		return Zipper{}, errors.New("no right child")
	}
	crumb := breadcrumb{
		parent: z.current,
		sibling: z.current.left,
		fromLeft: false,
	}
	return Zipper{
		current: z.current.right,
		trail: append(z.trail, crumb),
	}, nil
}

func (z Zipper) Up() (Zipper, error) {
	if len(z.trail) == 0 {
		return Zipper{}, errors.New("already at root")
	}
	crumb := z.trail[len(z.trail)-1]
	var parent *Node
	if crumb.fromLeft {
		if z.modified {
			parent = &Node{value: crumb.parent.value, left: z.current, right: crumb.sibling}
			return Zipper{
				current: parent,
				trail: z.trail[:len(z.trail)-1],
				modified: true,
			}, nil
		} else {
			return Zipper{
				current: crumb.parent,
				trail: z.trail[:len(z.trail)-1],
				modified: false,
			}, nil
		}
	} else {
		if z.modified {	
			parent = &Node{value: crumb.parent.value, left: crumb.sibling, right: z.current}
			return Zipper{
				current: parent,
				trail: z.trail[:len(z.trail)-1],
				modified: true,
			}, nil
		} else {
			return Zipper{
				current: crumb.parent,
				trail: z.trail[:len(z.trail)-1],
				modified: false,
			}, nil
		}
	}
}

func (z Zipper) SetValue(v int) Zipper {
	newNode := &Node{value: v, left: z.current.left, right: z.current.right}
	return Zipper{
		current: newNode,
		trail: z.trail,
		modified: true,
	}
}

func (z Zipper) SetLeft(v *Node) Zipper {
	newNode := &Node{value: z.current.value, left: v, right: z.current.right}
	return Zipper{
		current: newNode,
		trail: z.trail,
		modified: true,
	}
}

func (z Zipper) SetRight(v *Node) Zipper {
	newNode := &Node{value: z.current.value, left: z.current.left, right: v}
	return Zipper{
		current: newNode,
		trail: z.trail,
		modified: true,
	}
}
