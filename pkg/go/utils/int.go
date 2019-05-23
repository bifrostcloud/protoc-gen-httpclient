package utils

import "time"

// consts -
const (
	MaxUint = ^uint(0)
	MaxInt  = int(MaxUint >> 1)
)

// DurationToInt -
func DurationToInt(duration, unit time.Duration) int {
	durationAsNumber := duration / unit

	if int64(durationAsNumber) > int64(MaxInt) {
		return MaxInt
	}
	return int(durationAsNumber)
}
