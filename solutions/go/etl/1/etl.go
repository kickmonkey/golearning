package etl
import "strings"
func Transform(in map[int][]string) map[string]int {
    l := 0
    for _, v := range in {
        l += len(v)
    }
    a := make(map[string]int, l)
    for k, v := range in {
        for _, j := range v{
            a[strings.ToLower(j)] = k
        }
    }
    return a
}
