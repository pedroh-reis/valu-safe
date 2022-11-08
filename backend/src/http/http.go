package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pedroh-reis/valu-safe/backend/src/throw"
)

func sendResponse[T any](w http.ResponseWriter, data T, statusCode int) {
	response, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("content-type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "DELETE, POST, GET, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
	w.WriteHeader(statusCode)
	w.Write(response)

}

func sendResponseWithError(w http.ResponseWriter, serverError *throw.ServerError) {
	log.Print(serverError.Error())

	data := NewHTTPDefaultResponseBody(serverError.GetHttpError().GetMessage())
	sendResponse(w, data, serverError.GetHttpError().GetStatusCode())
}

func Exec[T any, K any](w http.ResponseWriter, r *http.Request, resolverFunc func(input T) (K, *throw.ServerError)) {
	input, err := DecodeBody[T](r.Body)
	if err != nil {
		sendResponseWithError(w, err)
		return
	}

	data, err := resolverFunc(input)
	if err != nil {
		sendResponseWithError(w, err)
		return
	}

	sendResponse(w, data, http.StatusOK)
}

func ExecWithoutInput[T any](w http.ResponseWriter, r *http.Request, resolverFunc func() (T, *throw.ServerError)) {
	data, err := resolverFunc()
	if err != nil {
		sendResponseWithError(w, err)
		return
	}

	sendResponse(w, data, http.StatusOK)
}
