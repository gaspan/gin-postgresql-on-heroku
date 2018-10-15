package model

// An Response struct represents http response to the client from the server
type Response struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    Employee `json:"data"`
}
