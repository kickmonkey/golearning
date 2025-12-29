package parsinglogfiles
import "regexp"
import "strings"

func IsValidLine(text string) bool {
    re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
    if text == "" {
        return []string{""}
    }
    re := regexp.MustCompile(`<[*~=\-]*>`) // * 允许空
    return re.Split(text, -1)
}


func CountQuotedPasswords(lines []string) int {
	// 找出每行中所有 "..." 片段（不处理转义引号的复杂情况，足够覆盖你给的用例）
	quoted := regexp.MustCompile(`"([^"]*)"`)

	count := 0
	for _, line := range lines {
		matches := quoted.FindAllStringSubmatch(line, -1)

		for _, m := range matches {
			insideQuotes := m[1] // 引号内部内容
			if strings.Contains(strings.ToLower(insideQuotes), "password") {
				count++
				break // 这一行已经算过了
			}
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
    re := regexp.MustCompile(`end-of-line\d*`)
    return re.ReplaceAllString(text, "") 
	panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User +([A-Za-z0-9]+)`)

	for i, line := range lines {
		// 查找是否匹配
		match := re.FindStringSubmatch(line)
		if len(match) == 2 { // 找到了用户名
			user := match[1]
			lines[i] = "[USR] " + user + " " + line
		}
	}
	return lines
}
            
