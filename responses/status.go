package responses

// Status is a struct
type Status struct {
	Status string `json:"status"`
}

type ErrorSplit struct {
	Error string `json:"error"`
}
