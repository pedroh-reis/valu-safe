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

func (I *Repository) IsLocked(input postgres.IsLockedInput) (bool, *throw.ServerError) {
	query := `
	SELECT locked
	FROM locker
	WHERE id = :id
	ORDER BY timestamp DESC
	LIMIT 1;
	`

	statement, err := I.db.PrepareNamed(query)
	if err != nil {
		return false, throw.NewServerError(err, ErrStmtNotCreated)
	}

	var isLocked bool
	err = statement.Get(&isLocked, input)
	// If there are no rows, the locker is locked
	if errors.Is(err, sql.ErrNoRows) {
		return true, nil
	} else if err != nil {
		return false, throw.NewServerError(err, ErrExecQuery)
	}

	return isLocked, nil
}

func (I *Repository) InsertState(input postgres.InsertStateInput) *throw.ServerError {
	query := `
	INSERT INTO locker (id, locked, timestamp)
	VALUES (:id, :locked, NOW());
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
