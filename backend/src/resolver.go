package main

import (
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
