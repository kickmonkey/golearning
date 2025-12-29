package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    sum := 0
    for _, v := range birdsPerDay{
        sum += v
    }
    return sum
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    sum := 0
    switch week{
    case 1:
    	for _, v := range birdsPerDay[0:7]{
        	sum += v
    	}
    	return sum
    case 2:
    	for _, v := range birdsPerDay[7:14]{
        	sum += v
    	}
    	return sum
    case 3:
    	for _, v := range birdsPerDay[14:21]{
        	sum += v
    	}
    	return sum
    
    	
	panic("Please implement the BirdsInWeek() function")
}
    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i, v := range birdsPerDay{
        if i == 0{
            birdsPerDay[i] = v+1
        }else if i > 0 && i % 2 == 0{
            birdsPerDay[i] = v+1
        }
    }
   return birdsPerDay
}
