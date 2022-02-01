package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Response struct {
	Data       interface{} `json : "data"`
	Message    string      `json : "message"`
	StatusCode int         `json : "statusCode"`
}

type ErrorResponse struct {
	StatusCode   int    `json:"status code"`
	ErrorMessage string `json:"error"`
}
