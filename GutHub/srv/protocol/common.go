package protocol

// ReturnData return handle result to client.
type ReturnData struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"` // in this case, return json only.
}
