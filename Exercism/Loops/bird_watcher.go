package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	dailyBirds := 0
	for i := 0; i < len(birdsPerDay); i++ {
		dailyBirds += birdsPerDay[i]
	}
	return dailyBirds
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weeklyBirds := 0
	for i := 0; i < len(birdsPerDay); i++ {
		if (len(birdsPerDay)%2) == 0 && len(birdsPerDay) > 7 {
			birdsPerDay = birdsPerDay[0 : week*7]
			weeklyBirds += birdsPerDay[i]
		} else if !((len(birdsPerDay) % 2) == 0) && len(birdsPerDay) > 7 {
			birdsPerDay = birdsPerDay[(week*7)-7 : (week * 7)] //[6:14] not counting index zero
			weeklyBirds += birdsPerDay[i]
		} else {
			birdsPerDay = birdsPerDay[0:]
			weeklyBirds += birdsPerDay[i]
		}
	}
	return weeklyBirds
	panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		if i%2 == 0 { // check if index number is even
			birdsPerDay[i] = birdsPerDay[i] + 1
		} else {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
	panic("Please implement the FixBirdCountLog() function")
}
