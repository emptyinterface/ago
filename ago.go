package ago

import (
	"log"
	"regexp"
	"strconv"
	"time"
)

var agoRegexp = regexp.MustCompile(`\s*([0-9,]+)\s+(yrs?|years?|mos?|months?|d|days?|h|hrs?|hours?|mins?|minutes?|s|secs?|seconds?|ms|millis?|milliseconds?|μs|micros?|microseconds?|nanos?|nanoseconds?),?\s+(?:ago)?\s*`)

func Parse(s string) (time.Time, error) {

	var d time.Duration

	for len(s) > 0 {
		ms := agoRegexp.FindAllStringSubmatch(s, 1)
		if len(ms) == 0 {
			break
		}
		s = s[len(ms[0][0]):]

		v, err := strconv.ParseFloat(ms[0][1], 64)
		if err != nil {
			return time.Time{}, err
		}
		switch ms[0][2] {
		case "yr", "yrs", "year", "years":
			d += time.Duration(v * float64(365*24*time.Hour)) // maybe jank
		case "mo", "mos", "month", "months":
			d += time.Duration(v * float64(30*24*time.Hour)) // maybe jank
		case "d", "day", "days":
			d += time.Duration(v * float64(24*time.Hour)) // maybe jank
		case "h", "hr", "hrs", "hour", "hours":
			d += time.Duration(v * float64(time.Hour))
		case "m", "min", "mins", "minute", "minutes":
			d += time.Duration(v * float64(time.Minute))
		case "s", "sec", "secs", "second", "seconds":
			d += time.Duration(v * float64(time.Second))
		case "ms", "milli", "millis", "millisecond", "milliseconds":
			d += time.Duration(v * float64(time.Millisecond))
		case "μs", "micro", "micros", "microsecond", "microseconds":
			d += time.Duration(v * float64(time.Microsecond))
		case "nano", "nanos", "nanosecond", "nanoseconds":
			d += time.Duration(v)
		default:
			log.Panicf("unrecognized interval: %q", ms[0][2])
		}
	}

	return time.Now().UTC().Add(-d), nil

}
