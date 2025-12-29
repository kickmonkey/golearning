package techpalace
import "fmt"
import "strings"
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
	panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    border := strings.Repeat("*", numStarsPerLine)
    return fmt.Sprintf("%s\n%s\n%s", border, welcomeMsg, border)
	panic("Please implement the AddBorder() function")
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    lines := strings.Split(strings.TrimSpace(oldMsg), "\n")
	if len(lines) < 2 {
		return strings.TrimSpace(oldMsg)
	}
	// 取中间行（这里假设第二行就是内容）
	mid := lines[1]
	return strings.TrimSpace(strings.Trim(mid, "*"))
	panic("Please implement the CleanupMessage() function")
}
