package tree

import (
  "errors"
  "sort"
)

// Define the Record type

type Record struct {
  ID     int
  Parent int
}

// Define the Node type

type Node struct {
  ID       int
  Children []*Node
}

func Build(records []Record) (*Node, error) {

  node := make(map[int]*Node, len(records))

  // Sort by ID
  sort.Slice(records, func(i, j int) bool {
    return records[i].ID < records[j].ID
  })

  for i, r := range records {
    if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
      return nil, errors.New("not in sequence or has bad parent")
    }
    node[r.ID] = &Node{ID: r.ID}
    if r.ID != 0 {
      node[r.Parent].Children = append(node[r.Parent].Children, node[r.ID])
    }
  }

  return node[0], nil
}
