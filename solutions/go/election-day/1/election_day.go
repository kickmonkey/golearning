package electionday
import "strconv"
// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
    count := &initialVotes
    return count
	//panic("Please implement the NewVoteCounter() function")
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
    if counter == nil {
        return 0
    }
    return *counter
	//panic("Please implement the VoteCount() function")
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int){
    *counter += increment
}


// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
    return &ElectionResult{
		Name: candidateName,
		Votes:         votes, 
    }
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
    return result.Name + " (" + strconv.Itoa(result.Votes) + ")"
	//panic("Please implement the DisplayResult() function")
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
    for k, v := range results{
        if k == candidate{
            v--
            results[candidate] = v
        }
    }
}
