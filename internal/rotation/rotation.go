// Package rotation provides stateless no-repeat selection for static content packs.
//
// The bot runs as a one-shot Fargate task with an ephemeral filesystem, so
// selection state cannot be persisted between runs. Instead of a random pick
// per day (which repeats items long before the pack is exhausted), PickIndex
// walks a deterministic shuffled order of the pack: day N of a cycle shows
// the Nth item of that cycle's permutation, and the pack is reshuffled only
// after every item has had its day. The same date always yields the same item.
package rotation

import (
	"math/rand"
	"time"
)

// epoch is an arbitrary fixed date all day counts are measured from.
var epoch = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)

// PickIndex returns the index to show for the given date seed (YYYYMMDD as an
// int, the same seed format every content pack already uses). Within any n
// consecutive days that fall in the same cycle, no index repeats.
func PickIndex(n, seed int) int {
	if n <= 1 {
		return 0
	}

	day := daysSinceEpoch(seed)
	cycle := day / n
	pos := day % n

	r := rand.New(rand.NewSource(int64(cycle)*1_000_003 + int64(n)))
	return r.Perm(n)[pos]
}

// WeekdaysSinceEpoch converts a YYYYMMDD seed into a count of weekdays
// (Mon-Fri) since the epoch. The bot only posts on weekdays, so rotating
// content types on this counter visits every type uniformly. Rotating on raw
// calendar days would alias with the week whenever the number of enabled
// types is a multiple of 7, pinning some types permanently to weekends.
func WeekdaysSinceEpoch(seed int) int {
	days := daysSinceEpoch(seed)
	count := (days / 7) * 5
	// epoch (2020-01-01) is a Wednesday
	const epochWeekday = 3
	for i := 0; i < days%7; i++ {
		wd := (epochWeekday + i) % 7 // 0=Sunday, 6=Saturday
		if wd != 0 && wd != 6 {
			count++
		}
	}
	return count
}

// daysSinceEpoch converts a YYYYMMDD seed into a day counter. time.Date
// normalizes out-of-range components, so synthetic test seeds still map to
// a stable day number.
func daysSinceEpoch(seed int) int {
	date := time.Date(seed/10000, time.Month((seed/100)%100), seed%100, 0, 0, 0, 0, time.UTC)
	days := int(date.Sub(epoch).Hours() / 24)
	if days < 0 {
		days = -days
	}
	return days
}
