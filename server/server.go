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

func Start(port int, path string, client *smtpclient.SmtpClient) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var reqBody requestBody
		err := json.NewDecoder(r.Body).Decode(&reqBody)
		if err != nil {
			writeFailureResponse(w, http.StatusBadRequest)
			return
		}
		emailErr := client.SendEmail(smtpclient.EmailRequest{
			Subject: reqBody.Subject,
			Body:    reqBody.Body,
		})
		if emailErr != nil {
			writeFailureResponse(w, http.StatusInternalServerError)
			return
		}
		writeSuccessResponse(w)
	})

	fmt.Printf("Starting server at %v:%v\n", path, port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		fmt.Printf("Error starting server:\n%v\n", err)
	}
}
