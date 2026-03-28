package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var total int = 0
    for _, count := range birdsPerDay {
        total += count
    }
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    m := make(map[int]int)
    
    for c := 0; c < len(birdsPerDay); c++ {
        if v, ok := m[c / 7]; ok {
            m[c / 7] = v + birdsPerDay[c]
        } else {
        	m[c / 7] = birdsPerDay[c]
        }
    }

	return m[week-1]
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i++ {
        if (i % 2 == 0) {
            birdsPerDay[i] = birdsPerDay[i] + 1
        }
    }
	return birdsPerDay
}
