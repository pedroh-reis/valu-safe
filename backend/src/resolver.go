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

	state, err := I.GetState(GetStateInput{Id: input.Id})
	if err != nil {
		return nil, err
	}

	err = I.repository.InsertState(postgres.InsertStateInput{
		Id:       input.Id,
		Unlocked: !state.IsUnlocked,
	})
	if err != nil {
		return nil, err
	}

	return http.NewHTTPDefaultResponseBody(result.Body.Message), nil
}

var ErrIncorrectTimeframe = throw.NewHttpError("Invalid timeframe. Check the documentation.", goHttp.StatusBadRequest)

func (I *Resolver) GetState(input GetStateInput) (*GetStateResult, *throw.ServerError) {
	if input.Timestamp.IsZero() {
		input.Timestamp = time.Now().UTC()
	}

	isUnlocked, err := I.repository.IsUnlocked(postgres.IsUnlockedInput{
		Id:        input.Id,
		Timestamp: input.Timestamp,
	})
	if err != nil {
		return nil, err
	}

	return &GetStateResult{
		IsUnlocked: isUnlocked,
	}, nil
}

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

	state, err := I.GetState(GetStateInput{Id: input.Id, Timestamp: timestamp})
	if err != nil {
		return nil, err
	}

	history, err := I.repository.GetHistory(postgres.GetHistoryInput{Id: input.Id, Timestamp: timestamp})
	if err != nil {
		return nil, err
	}

	firstHistory := postgres.GetHistoryResult{Unlocked: state.IsUnlocked, Timestamp: timestamp}
	history = append([]postgres.GetHistoryResult{firstHistory}, history...)

	timesUnlocked := len(history) / 2
	if history[0].Unlocked && len(history)%2 == 1 {
		timesUnlocked += 1
	}

	var totalTime int64

	for i := 0; i < len(history)-1; i++ {
		totalTime += history[i+1].Timestamp.Unix() - history[i].Timestamp.Unix()
	}

	var percUnlocked float32
	var percLocked float32

	if len(history) == 1 {
		if history[0].Unlocked {
			percUnlocked = 100.0
		} else {
			percUnlocked = 0.0
		}
		percLocked = 100.0 - percUnlocked
	} else {
		if history[0].Unlocked {
			percUnlocked = 100 * float32(totalTime) / float32(currentTime.Unix()-timestamp.Unix())
			percLocked = 100.0 - percUnlocked
		} else {
			percLocked = 100 * float32(totalTime) / float32(currentTime.Unix()-timestamp.Unix())
			percUnlocked = 100.0 - percLocked
		}
	}

	return &GetStatisticsResult{
		TimesUnlocked: timesUnlocked,
		PercUnlocked:  percUnlocked,
		PercLocked:    percLocked,
		History:       history,
	}, nil
}
