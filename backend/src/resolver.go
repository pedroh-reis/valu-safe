package main

import (
	goHttp "net/http"
	"time"

	"github.com/pedroh-reis/valu-safe/backend/src/client"
	"github.com/pedroh-reis/valu-safe/backend/src/http"
	"github.com/pedroh-reis/valu-safe/backend/src/postgres"
	"github.com/pedroh-reis/valu-safe/backend/src/throw"
)

type Resolver struct {
	repository *Repository
	client     *client.Client
}

func NewResolver(repository *Repository) *Resolver {
	return &Resolver{
		repository: repository,
		client:     client.NewClient(),
	}
}

func (I *Resolver) Home() (*http.HTTPDefaultResponseBody, *throw.ServerError) {
	return http.NewHTTPDefaultResponseBody("Welcome to ValuSafe Backend!"), nil
}

func (I *Resolver) ChangeState(input ChangeStateInput) (*http.HTTPDefaultResponseBody, *throw.ServerError) {
	result, err := I.client.ChangeState(input.Id)
	if err != nil {
		return nil, err
	} else if result.StatusCode != 200 {
		httpError := throw.NewHttpError(result.Body.Message, result.StatusCode)
		return nil, throw.NewServerError(nil, httpError)
	}

	isLocked, err := I.repository.IsLocked(postgres.IsLockedInput{
		Id: input.Id,
	})
	if err != nil {
		return nil, err
	}

	err = I.repository.InsertState(postgres.InsertStateInput{
		Id:     input.Id,
		Locked: !isLocked,
	})
	if err != nil {
		return nil, err
	}

	return http.NewHTTPDefaultResponseBody(result.Body.Message), nil
}

var ErrIncorrectTimeframe = throw.NewHttpError("Invalid timeframe. Check the documentation.", goHttp.StatusBadRequest)

func (I *Resolver) GetStatistics(input GetStatisticsInput) (*GetStatisticsResult, *throw.ServerError) {
	var timestamp time.Time
	currentTime := time.Now().UTC()
	switch input.Timeframe {
	case "second":
		timestamp = currentTime.Add(-time.Second * time.Duration(input.Value))
	case "minute":
		timestamp = currentTime.Add(-time.Minute * time.Duration(input.Value))
	case "hour":
		timestamp = currentTime.Add(-time.Hour * time.Duration(input.Value))
	case "day":
		timestamp = currentTime.AddDate(0, 0, -input.Value)
	case "month":
		timestamp = currentTime.AddDate(0, -input.Value, 0)
	case "year":
		timestamp = currentTime.AddDate(-input.Value, 0, 0)
	default:
		return nil, throw.NewServerError(nil, ErrIncorrectTimeframe)
	}

	history, err := I.repository.GetHistory(postgres.GetHistoryInput{
		Id:        input.Id,
		Timestamp: timestamp,
	})
	if err != nil {
		return nil, err
	}

	timeUnlocked := 0
	if len(history) > 0 {
		timeUnlocked = len(history) / 2
		if !history[0].Locked && len(history)%2 == 1 {
			timeUnlocked += 1
		}
	}

	return &GetStatisticsResult{
		TimesUnlocked: timeUnlocked,
		History:       history,
	}, nil
}
