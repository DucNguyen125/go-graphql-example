package convert

import "time"

func TimeToString(t time.Time) *string {
	stringTime := t.Format(time.RFC3339)
	return &stringTime
}
