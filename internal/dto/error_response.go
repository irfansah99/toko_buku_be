package dto


type ErrorResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Code    string            `json:"code,omitempty"`  
	Errors  map[string]string `json:"errors,omitempty"`
	Data    interface{}       `json:"data,omitempty"`   
}