package dto

type BaseResponse struct {
	Err  string      `json:"error"`
	Data interface{} `json:"data"`
}
