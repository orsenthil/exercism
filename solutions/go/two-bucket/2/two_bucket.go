package twobucket

import "errors"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type bucket struct {
	name  string
	size  int
	level int
}

func bucketFromValues(name string, size int) *bucket {
	b := bucket{name: name, size: size, level: 0}
	return &b
}

func (buc bucket) room() int {
	return buc.size - buc.level
}

func (buc bucket) full() bool {
	return buc.level == buc.size
}

func (buc bucket) isEmpty() bool {
	return buc.level == 0
}

func (buc *bucket) fill() {
	buc.level = buc.size
}

func (buc *bucket) empty() {
	buc.level = 0
}

func (buc *bucket) pourInto(other *bucket) {
	amount := min(buc.level, other.room())
	buc.level -= amount
	other.level += amount
}

type resolved struct {
	status string
	target bucket
	other  bucket
}

func hitTarget(goal int, start bucket, other bucket) resolved {
	if start.level == goal {
		return resolved{"resolved", start, other}
	}
	if other.level == goal {
		return resolved{"resolved", other, start}
	}
	return resolved{"not", start, other}
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	if goalAmount > max(sizeBucketOne, sizeBucketTwo) || goalAmount < 1 {
		return "None", sizeBucketOne, sizeBucketTwo, errors.New("Invalid goal amount")
	}
	if min(sizeBucketOne, sizeBucketTwo) < 1 {
		return "None", sizeBucketOne, sizeBucketTwo, errors.New("Invalid bucket size")
	}
	var start *bucket
	if startBucket == "one" {
		start = bucketFromValues("one", sizeBucketOne)
	} else {
		start = bucketFromValues("two", sizeBucketTwo)
	}
	var other *bucket
	if startBucket == "two" {
		other = bucketFromValues("one", sizeBucketOne)
	} else {
		other = bucketFromValues("two", sizeBucketTwo)
	}

	moves := 0

	if other.size == goalAmount {
		return other.name, 2, start.size, nil
	}

	for {
		resolution := hitTarget(goalAmount, *start, *other)
		if resolution.status == "not" {
			if start.isEmpty() {
				start.fill()
			} else if other.full() {
				other.empty()
			} else {
				start.pourInto(other)
			}

			moves += 1

			if start.full() && other.full() {
				return "None", sizeBucketOne, sizeBucketTwo, errors.New("No solution")
			}
		} else {
			return resolution.target.name, moves, resolution.other.level, nil
		}

	}

}
