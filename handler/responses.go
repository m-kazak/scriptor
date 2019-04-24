package handler

//GeneralResponse
type GeneralResponse struct {
	ErrorCode   	int 	`json:"err_code"`
	ErrorMessage    string 	`json:"err_msg"`
}

type LoginResponse struct {
	ErrorCode   	int 	`json:"err_code"`
	ErrorMessage    string 	`json:"err_msg"`
	Name			string  `json:"name"`
	Email			string  `json:"email"`
	Session			string  `json:"session"`
}