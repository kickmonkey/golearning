package bob

import "strings"

func Hey(remark string) string {
    remark = strings.TrimSpace(remark)

    if remark == "" {
        return "Fine. Be that way!"
    }

    isQuestion := strings.HasSuffix(remark, "?")

    // shouting: 全大写 + 至少有一个字母
    hasLetter := false
    for _, r := range remark {
        if 'a' <= r && r <= 'z' || 'A' <= r && r <= 'Z' {
            hasLetter = true
            break
        }
    }
    isShouting := hasLetter && remark == strings.ToUpper(remark)

    if isShouting && isQuestion {
        return "Calm down, I know what I'm doing!"
    }
    if isShouting {
        return "Whoa, chill out!"
    }
    if isQuestion {
        return "Sure."
    }
    return "Whatever."
}