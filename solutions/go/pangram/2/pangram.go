package pangram
import "strings"

func IsPangram(input string) bool {
    count := 0
    input = strings.ToLower(input)
    al := "abcdefghijklmnopqrstuvwxyz"
    for _, v := range al{
        if strings.ContainsRune(input, v) {
            count++
		}
    }
    if count == 26 {
        return true
    }
    return false
}
