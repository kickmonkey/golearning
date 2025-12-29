package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
    c := 0
    v, ok := cb[file]
    if !ok{
        return 0
    }else{
        for _, j := range v{
            if j == true {
                c++
            }
        }
    }
    return c
	panic("Please implement CountInFile()")
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
    c := 0
    if rank > 8 || rank < 1{
        return 0
    }
    for _, v := range cb{
        for i, vv := range v{
            if i == rank-1 && vv == true{
                c++
            }
        }
    }
    return c
	panic("Please implement CountInRank()")
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
    total := 0
    for _, file := range cb {
        total += len(file)
    }
    return total

}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    c := 0
    for _, v := range cb{
        for _, j := range v{
            if j == true {
                c++
            }
        }
    }
    return c
	panic("Please implement CountOccupied()")
}
