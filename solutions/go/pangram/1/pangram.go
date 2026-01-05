package pangram
import "strings"

func IsPangram(input string) bool {
    count := 0
    al := "abcdefghijklmnopqrstuvwxyz"
    for _, v := range al{
        if strings.ContainsRune(strings.ToLower(input), v) {
            count++
		}
    }
    if count == 26 {
        return true
    }
    return false
	panic("Please implement the IsPangram function")
}
