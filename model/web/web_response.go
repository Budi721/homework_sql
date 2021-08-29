package web

type WebResponse struct {
	Code   int         `json:"code"`
	Error  interface{} `json:"error"`
	Result interface{} `json:"result"`
}
