package time

import "time"

func GetNow() string {
	return time.Now().Format(time.RFC3339)
}

func GetNowUTC() string {
	return time.Now().UTC().Format(time.RFC3339)
}
