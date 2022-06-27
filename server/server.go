package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/david-allan-jones/contactserver/smtpclient"
)

func writeFailureResponse(w http.ResponseWriter, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(responseBody{false})
}

func writeSuccessResponse(w http.ResponseWriter) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(responseBody{true})
}

func Start(config ServerConfig) {
	http.HandleFunc(config.Path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqBody requestBody
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			writeFailureResponse(w, http.StatusBadRequest)
			return
		}
		emailErr := config.SmtpClient.SendEmail(smtpclient.EmailRequest{
			Name:    reqBody.Name,
			Contact: reqBody.Contact,
			Message: reqBody.Message,
		})
		if emailErr != nil {
			writeFailureResponse(w, http.StatusInternalServerError)
			return
		}
		writeSuccessResponse(w)
	})

	fmt.Printf("Starting server at %v:%v\n", config.Path, config.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", config.Port), nil)
	if err != nil {
		fmt.Printf("Error starting server:\n%v\n", err)
	}
}
