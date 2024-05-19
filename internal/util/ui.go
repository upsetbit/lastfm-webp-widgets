package util

import (
	"fmt"
	"math"
	"time"
)

const (
	HOURS_IN_DAY   = 24 * time.Hour
	HOURS_IN_MONTH = 30 * HOURS_IN_DAY
	HOURS_IN_YEAR  = 365 * HOURS_IN_DAY
)

func durationToHumanReadable(duration time.Duration, durationType time.Duration, word string) string {
	dur := int(math.Ceil(duration.Hours() / durationType.Hours()))
	if dur > 1 {
		word += "s"
	}

	return fmt.Sprintf("%d %s", dur, word)
}

func UnixToRelativeHumanTime(unix int64) string {
	diff := time.Since(time.Unix(unix, 0))

	if diff >= HOURS_IN_YEAR {
		return "more than a year"
	}

	if diff >= HOURS_IN_MONTH {
		return durationToHumanReadable(diff, HOURS_IN_MONTH, "month")
	}

	if diff >= HOURS_IN_DAY {
		return durationToHumanReadable(diff, HOURS_IN_DAY, "day")
	}

	if diff >= time.Hour {
		return durationToHumanReadable(diff, time.Hour, "hour")
	}

	if diff >= time.Minute {
		return durationToHumanReadable(diff, time.Minute, "minute")
	}

	return durationToHumanReadable(diff, time.Second, "second")
}

func UnixToHumanReadable(unix int64) string {
	return time.Unix(unix, 0).Format("02 Jan 2006")
}

func BoolToYesOrNo(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
