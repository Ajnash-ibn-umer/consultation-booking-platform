package common

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Errors  string      `json:"errors,omitempty"`
}
