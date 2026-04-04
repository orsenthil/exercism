package satellite

import "errors"

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func TreeFromTraversals(preorder, inorder []string) (*Node, error) {
	if len(preorder) != len(inorder) {
		return nil, errors.New("traversals must have the same length")
	}
	if len(preorder) == 0  && len(inorder) == 0 {
		return nil, nil
	}
	if len(preorder) == 1 && len(inorder) == 1 {
		return &Node{Value: preorder[0]}, nil
	}
	if !verifyUnique(preorder) || !verifyUnique(inorder) {
		return nil, errors.New("traversals must contain unique items")
	}

	root := preorder[0]

	rootIndex := indexOf(root, inorder)
	if rootIndex == -1 {
		return nil, errors.New("traversals must have the same elements")
	}
	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	leftPreorder := preorder[1:len(leftInorder)+1]
	rightPreorder := preorder[len(leftInorder)+1:]


	left, err := TreeFromTraversals(leftPreorder, leftInorder)
	if err != nil {
		return nil, err
	}
	right, err := TreeFromTraversals(rightPreorder, rightInorder)
	if err != nil {
		return nil, err
	}
	return &Node{Value: root, Left: left, Right: right}, nil

}


func indexOf(value string, array []string) int {
	for i, v := range array {
		if v == value {
			return i
		}
	}
	// should never happen
	return -1
}

func verifyUnique(array []string) bool {
	seen := make(map[string]bool)
	for _, v := range array {
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}