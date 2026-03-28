package raindrops

import (
    "strconv"
    "strings"
    )
    

func Convert(number int) string {
    var sb strings.Builder
    var factor bool = false

    if (number % 3 == 0) {
        factor = true
        sb.WriteString("Pling")
    }

    if (number % 5 == 0) {
        factor = true
        sb.WriteString("Plang")
    }

    if (number % 7 == 0) {
        factor = true
        sb.WriteString("Plong")
    }

    if (factor == false) {
        return strconv.Itoa(number)
    }

    return sb.String()
}
