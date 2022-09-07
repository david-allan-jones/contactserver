package server

type requestBody struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Message string `json:"message"`
}
