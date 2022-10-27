package client

import (
	"fmt"
	"net/http"

	"github.com/pedroh-reis/valu-safe/backend/src/config"
	utilHttp "github.com/pedroh-reis/valu-safe/backend/src/http"
	"github.com/pedroh-reis/valu-safe/backend/src/throw"
)

var ErrClientRequest = throw.NewHttpError("Client could not execute a request", http.StatusInternalServerError)

type Client struct {
	config     config.MicrocontrollerServerConfig
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		config:     config.NewMicrocontrollerServerConfig(),
		httpClient: http.DefaultClient,
	}
}

func (I *Client) ChangeState(id string) (*ChangeStateResult, *throw.ServerError) {
	endpoint := fmt.Sprintf("/locker/%s", id)
	url := fmt.Sprintf("http://%s%s", I.config.GetAddress(), endpoint)

	response, err := http.Get(url)
	if err != nil {
		return nil, throw.NewServerError(err, ErrClientRequest)
	}

	responseBody, serverError := utilHttp.DecodeBody[ChangeStateBodyResult](response.Body)
	if serverError != nil {
		return nil, serverError
	}

	return &ChangeStateResult{
		Body:       responseBody,
		StatusCode: response.StatusCode,
	}, nil
}
