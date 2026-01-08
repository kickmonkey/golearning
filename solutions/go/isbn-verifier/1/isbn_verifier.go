package isbn

import "strings"

func IsValidISBN(isbn string) bool {
    isbn = strings.ReplaceAll(isbn, "-", "")
    if len(isbn) != 10 {
        return false
    }

    sum := 0
    for i := 0; i < 9; i++ {
        c := isbn[i]
        if c < '0' || c > '9' {
            return false
        }
        sum += int(c-'0') * (10 - i)
    }

    last := isbn[9]
    if last == 'X' {
        sum += 10
    } else {
        if last < '0' || last > '9' {
            return false
        }
        sum += int(last - '0')
    }

    return sum%11 == 0
}