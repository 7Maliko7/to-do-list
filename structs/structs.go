package structs

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
