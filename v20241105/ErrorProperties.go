package v20241105

type ErrorProperties struct {
	Code    Cursor `json:"code"`
	Data    Data   `json:"data"`
	Message Cursor `json:"message"`
}
