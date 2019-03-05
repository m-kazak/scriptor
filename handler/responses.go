package handler

//GeneralResponse
type GeneralResponse struct {
	ErrorCode   	int 	`json:"err_code"`
	ErrorMessage    string 	`json:"err_msg"`
}