package common

import "net/http"

type Resp struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Errors  error       `json:"errors,omitempty"`
	Success bool        `json:"success"`
}

type S struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

type E struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Errors  error  `json:"errors,omitempty"`
}

func NewSuccessResponse(r *S) *Resp {
	if r.Message == "" {
		r.Message = "Api Processing completed"
	}
	if r.Status == 0 {
		r.Status = http.StatusOK
	}
	return &Resp{
		Message: r.Message,
		Status:  r.Status,
		Data:    r.Data,
		Errors:  nil,
		Success: true,
	}
}

func NewErrorResponse(r *E) *Resp {
	if r.Message == "" {
		r.Message = "Something went wrong"
	}
	if r.Status == 0 {
		r.Status = http.StatusBadGateway
	}
	return &Resp{
		Message: r.Message,
		Status:  r.Status,
		Data:    nil,
		Errors:  r.Errors,
		Success: false,
	}
}
