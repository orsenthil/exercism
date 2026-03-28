package tree

// Define the Record type

type Record struct {
	id       int
	parentId int
}

// Define the Node type

type Node struct {
	record   Record
	children []Record
}

func Build(records []Record) (*Node, error) {
	panic("Please implement the Build function")
}
