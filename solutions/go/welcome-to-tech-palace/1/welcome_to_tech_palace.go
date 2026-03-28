package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    var result string
    result = strings.Repeat("*", numStarsPerLine) + "\n"
    result = result + welcomeMsg + "\n"
    result = result + strings.Repeat("*", numStarsPerLine)
    return result
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    var result string
    result = strings.ReplaceAll(oldMsg, "*", "")
    result = strings.ReplaceAll(result, "\n", "")
    result = strings.TrimSpace(result)
    return result
}
