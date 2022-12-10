package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pedroh-reis/valu-safe/backend/src/throw"
)

var ErrDecodeBody = throw.NewHttpError("Could not decode body as json", http.StatusBadRequest)
var ErrEncodeBody = throw.NewHttpError("Could not encode body as json", http.StatusBadRequest)

var ErrDecodeBodyAsByte = throw.NewHttpError("Could not decode body in bytes", http.StatusBadRequest)

func DecodeBody[T any](body io.ReadCloser) (T, *throw.ServerError) {
	var structBody T

	err := json.NewDecoder(body).Decode(&structBody)
	if err != nil {
		return structBody, throw.NewServerError(err, ErrDecodeBody)
	}

	return structBody, nil
}

func EncodeBody[T any](input T) (*bytes.Buffer, *throw.ServerError) {
	body := new(bytes.Buffer)

	err := json.NewEncoder(body).Encode(input)
	if err != nil {
		return nil, throw.NewServerError(err, ErrEncodeBody)
	}

	return body, nil
}
