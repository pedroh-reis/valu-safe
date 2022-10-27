package postgres

import "time"

type Locker struct {
	Id        int       `db:"id"`
	Locked    bool      `db:"locked"`
	Timestamp time.Time `db:"timestamp"`
}

type InsertStateInput struct {
	Id     string `json:"id"`
	Locked bool   `json:"locked"`
}

type IsLockedInput struct {
	Id string `json:"id"`
}

type GetHistoryInput struct {
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

type GetHistoryResult struct {
	Locked    bool      `json:"locked"`
	Timestamp time.Time `json:"timestamp"`
}
