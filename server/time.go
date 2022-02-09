package server

import "time"

func TimeNow() time.Time {
	return time.Now().UTC()
}

func FormatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
