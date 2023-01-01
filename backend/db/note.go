package db

import "time"

// TODO
type Note struct {
	// note's url
	ID        string
	CreatedAt time.Time

	InReplyToID string
	Pubished    time.Time
}
