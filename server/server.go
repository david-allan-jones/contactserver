package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/david-allan-jones/contactserver/smtpclient"
)

func getPort(useHttps bool) int {
	if useHttps {
		return 443
	}
	return 80
}

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

	port := getPort(config.UseHttps)
	fmt.Printf("Starting server at %v:%v\n", config.Path, port)
	portStr := fmt.Sprintf(":%v", port)
	if config.UseHttps {
		err := http.ListenAndServeTLS(portStr, config.TlsCert, config.TlsKey, nil)
		if err != nil {
			fmt.Printf("Error starting HTTPS server:\n%v\n", err)
		}
	} else {
		err := http.ListenAndServe(portStr, nil)
		if err != nil {
			fmt.Printf("Error starting HTTP server:\n%v\n", err)
		}
	}
}
