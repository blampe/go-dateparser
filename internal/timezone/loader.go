package timezone

import (
	"fmt"
	"strings"
	"time"
)

func Load(name string) (*time.Location, error) {
	// Check if name exists in known timezone list
	upperName := strings.ToUpper(name)
	for _, info := range timezoneInfoList {
		if offset, exist := info.Timezones[upperName]; exist {
			return time.FixedZone(name, offset), nil
		}
	}

	// Try to look up from IANA Time Zone database
	loc, err := time.LoadLocation(name)
	if err == nil && loc != nil {
		return loc, nil
	}

	// Try to pop timezone
	_, tz := PopTzOffset(" " + name)
	if !tz.IsZero() {
		return time.FixedZone(name, tz.Offset), nil
	}

	return time.UTC, fmt.Errorf("%s is not a known timezone", name)
}

func MustLoad(name string) *time.Location {
	loc, err := Load(name)
	if loc == nil || err != nil {
		return time.UTC
	}
	return loc
}
