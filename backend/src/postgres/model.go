package postgres

import "time"

type Locker struct {
	Id        int       `db:"id"`
	Locked    bool      `db:"locked"`
	Timestamp time.Time `db:"timestamp"`
}

type InsertStateInput struct {
	Id       string `json:"id"`
	Unlocked bool   `json:"unlocked"`
}

type IsUnlockedInput struct {
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

type GetHistoryInput struct {
	Id        string    `json:"id"`
	Timestamp time.Time `json:"timestamp"`
}

type GetHistoryResult struct {
	Unlocked  bool      `json:"unlocked"`
	Timestamp time.Time `json:"timestamp"`
}
