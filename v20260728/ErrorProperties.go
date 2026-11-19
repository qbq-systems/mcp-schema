package v20260728

type ErrorProperties struct {
	Code    Cursor            `json:"code"`
	Data    StructuredContent `json:"data"`
	Message Cursor            `json:"message"`
}
