package relativedistance

type entry struct {
	name string
	distance int
}

func DegreeOfSeparation(familyTree map[string][]string, personA, personB string) (int, bool) {
	biFamilyTree := biFamilyTree(familyTree)

	queue := []entry{{name: personA, distance: 0}}
	visited := make(map[string]bool)
	visited[personA] = true


	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.name == personB {
			return current.distance, true
		}
		for _, neighbor := range biFamilyTree[current.name] {
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			queue = append(queue, entry{name: neighbor, distance: current.distance + 1})
		}
	}
	return 0, false
}

func biFamilyTree(familyTree map[string][]string) map[string][]string {

	newFamilyTree := make(map[string][]string)
	for person, children := range familyTree {
		newFamilyTree[person] = append(newFamilyTree[person], children...)
		for _, child := range children {
			newFamilyTree[child] = append(newFamilyTree[child], person)
			for _, sibling := range children {
				if sibling != child {
					newFamilyTree[child] = append(newFamilyTree[child], sibling)
				}
			}
		}
	}
	return newFamilyTree

}


