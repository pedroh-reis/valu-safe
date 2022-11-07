package main

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/pedroh-reis/valu-safe/backend/src/postgres"
	"github.com/pedroh-reis/valu-safe/backend/src/throw"
)

var ErrStmtNotCreated = throw.NewHttpError("Could not created SQL named statement", http.StatusInternalServerError)
var ErrExecQuery = throw.NewHttpError("Could not execute query", http.StatusInternalServerError)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (I *Repository) IsUnlocked(input postgres.IsUnlockedInput) (bool, *throw.ServerError) {
	query := `
	SELECT unlocked
	FROM locker
	WHERE id = :id AND timestamp < :timestamp
	ORDER BY timestamp DESC
	LIMIT 1;
	`

	statement, err := I.db.PrepareNamed(query)
	if err != nil {
		return false, throw.NewServerError(err, ErrStmtNotCreated)
	}

	var isUnlocked bool
	err = statement.Get(&isUnlocked, input)
	// If there are no rows, the locker is locked
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	} else if err != nil {
		return false, throw.NewServerError(err, ErrExecQuery)
	}

	return isUnlocked, nil
}

func (I *Repository) InsertState(input postgres.InsertStateInput) *throw.ServerError {
	query := `
	INSERT INTO locker (id, unlocked, timestamp)
	VALUES (:id, :unlocked, NOW());
	`

	statement, err := I.db.PrepareNamed(query)
	if err != nil {
		return throw.NewServerError(err, ErrStmtNotCreated)
	}

	_, err = statement.Exec(input)
	if err != nil {
		return throw.NewServerError(err, ErrExecQuery)
	}

	return nil
}

func (I *Repository) GetHistory(input postgres.GetHistoryInput) ([]postgres.GetHistoryResult, *throw.ServerError) {
	query := `
	SELECT unlocked, timestamp
	FROM locker
	WHERE id = :id AND timestamp >= :timestamp
	ORDER BY timestamp;
	`

	statement, err := I.db.PrepareNamed(query)
	if err != nil {
		return nil, throw.NewServerError(err, ErrStmtNotCreated)
	}

	var result []postgres.GetHistoryResult
	err = statement.Select(&result, input)
	if err != nil {
		return nil, throw.NewServerError(err, ErrExecQuery)
	}

	if len(result) == 0 {
		return []postgres.GetHistoryResult{}, nil
	}

	return result, nil
}
