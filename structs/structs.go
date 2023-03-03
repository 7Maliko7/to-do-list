package structs

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type OkResponse struct {
	Code int `json:"code"`
}
