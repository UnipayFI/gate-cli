package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const dateTimeLayout = "2006-01-02 15:04:05"

// FormatUnixTime renders an int64 timestamp as a UTC string. The unit (s / ms /
// μs / ns) is sniffed by magnitude so that endpoints returning different units
// for the same conceptual field all render correctly. Zero or negative values
// produce an empty string.
func FormatUnixTime(v int64) string {
	if v <= 0 {
		return ""
	}
	var t time.Time
	switch {
	case v >= 1e18:
		t = time.Unix(0, v)
	case v >= 1e15:
		t = time.UnixMicro(v)
	case v >= 1e12:
		t = time.UnixMilli(v)
	default:
		t = time.Unix(v, 0)
	}
	return t.UTC().Format(dateTimeLayout)
}

// FormatUnixSeconds renders a Unix-seconds timestamp (Gate's REST APIs report
// most times in whole or fractional seconds) as a UTC string. Zero or negative
// values produce an empty string.
func FormatUnixSeconds(v float64) string {
	if v <= 0 {
		return ""
	}
	sec := int64(v)
	nsec := int64((v - float64(sec)) * 1e9)
	return time.Unix(sec, nsec).UTC().Format(dateTimeLayout)
}

// FormatTime renders a time.Time as a UTC string in the project layout.
func FormatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format(dateTimeLayout)
}

// ParseTimeFlag accepts either a Unix-seconds timestamp or a
// "YYYY-MM-DD HH:MM:SS" datetime (interpreted in local time) and returns the
// corresponding Unix-seconds value. An empty or "0" input yields (0, false, nil).
func ParseTimeFlag(flagName string, value string) (int64, bool, error) {
	v := strings.TrimSpace(value)
	if v == "" || v == "0" {
		return 0, false, nil
	}

	if isAllDigits(v) {
		sec, err := strconv.ParseInt(v, 10, 64)
		if err != nil || sec < 0 {
			return 0, false, invalidTimeFlagError(flagName)
		}
		return sec, true, nil
	}

	t, err := time.ParseInLocation(dateTimeLayout, v, time.Local)
	if err != nil {
		return 0, false, invalidTimeFlagError(flagName)
	}
	return t.Unix(), true, nil
}

func invalidTimeFlagError(flagName string) error {
	name := strings.TrimSpace(flagName)
	if name == "" {
		name = "time"
	}
	return fmt.Errorf("invalid %s: expected unix seconds timestamp (e.g. 1734495381) or datetime \"YYYY-MM-DD HH:MM:SS\" (e.g. \"2025-12-18 04:16:21\")", name)
}

func isAllDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
