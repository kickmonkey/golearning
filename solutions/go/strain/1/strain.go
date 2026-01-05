package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any](list []T, predicate func(T) bool) []T{
    out := make([]T, 0, len(list))
    for _, v := range list {
        if predicate(v){
           out = append(out, v)
        }
}
    return out
}

func Discard[T any](list []T, predicate func(T) bool) []T {
    return Keep(list, func(v T) bool { return !predicate(v) })
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
