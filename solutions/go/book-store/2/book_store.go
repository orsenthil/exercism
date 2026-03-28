// Package bookstore provides functionality for calculating discounts for combinations of books.
package bookstore
import (
	"sort"
)
// Cost calculates the total price of the given books with the biggest discount.
func Cost(books []int) int {
	bookMap := map[int]int{1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
	titleCounts := make([]int, 5, 5)
	for i := 0; i < len(books); i++ {
		bookMap[books[i]]++
	}
	i := 0
	for _, v := range bookMap {
		titleCounts[i] = v
		i++
	}
	sort.Ints(titleCounts)
	return costAll(titleCounts)
}
func costTwo(titleCounts []int) int {
	return 2*760*titleCounts[0] + 800*(titleCounts[1]-titleCounts[0])
}
func costThree(titleCounts []int) int {
	return 3*720*titleCounts[0] + costTwo([]int{titleCounts[1] - titleCounts[0], titleCounts[2] - titleCounts[0]})
}
func costFour(titleCounts []int) int {
	return 4*640*titleCounts[0] + costThree([]int{titleCounts[1] - titleCounts[0], titleCounts[2] - titleCounts[0], titleCounts[3] - titleCounts[0]})
}
func costAll(titleCounts []int) int {
	a, b, c, d, e := titleCounts[0], titleCounts[1], titleCounts[2], titleCounts[3], titleCounts[4]
	// Compare two methods and choose the cheaper one.
	// Method 1: Prefer Groups of Five:
	price := 5*600*a + costFour([]int{b - a, c - a, d - a, e - a})
	if a == 0 {
		// return immediately since Method 1 and Method 2 give the same result with 4 different titles or less.
		return price
	}
	// Method 2: Prefer Groups of Four:
	subtractFirst := true
	groupOfFour := 0
	var priceAlternative int
	for a != 0 && a != e {
		groupOfFour++
		if subtractFirst {
			a--
		} else {
			b--
		}
		c, d, e = c-1, d-1, e-1
		subtractFirst = !subtractFirst
	}
	if a != 0 {
		priceAlternative = 4*640*groupOfFour + 5*600*a
	} else {
		newTitleCounts := []int{b, c, d, e}
		sort.Ints(newTitleCounts)
		priceAlternative = 4*640*groupOfFour + costFour(newTitleCounts)
	}
	if priceAlternative < price {
		price = priceAlternative
	}
	return price
}
