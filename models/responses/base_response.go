package responses

type SuccessBaseResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorBaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
