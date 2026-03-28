package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	var len1, len2 = len(l1), len(l2)
	if len1 == 0 && len2 == 0 {
		return RelationEqual
	}
	if len1 == 0 && len2 > 0 {
		return RelationSublist
	}
	if len1 > 0 && len2 == 0 {
		return RelationSuperlist
	}
	if len1 == len2 {
		for i := 0; i < len1; i++ {
			if l1[i] != l2[i] {
				return RelationUnequal
			}
		}
		return RelationEqual
	}
	if len1 < len2 {
		for i := 0; i < len2-len1+1; i++ {
			if l1[0] == l2[i] {
				for j := 0; j < len1; j++ {
					if l1[j] != l2[i+j] {
						break
					}
					if j == len1-1 {
						return RelationSublist
					}
				}
			}
		}
	}
	if len1 > len2 {
		for i := 0; i < len1-len2+1; i++ {
			if l2[0] == l1[i] {
				for j := 0; j < len2; j++ {
					if l2[j] != l1[i+j] {
						break
					}
					if j == len2-1 {
						return RelationSuperlist
					}
				}
			}
		}
	}
	return RelationUnequal
}
