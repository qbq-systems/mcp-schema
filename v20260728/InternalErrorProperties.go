package v20260728

type InternalErrorProperties struct {
	Code    FluffyCode        `json:"code"`
	Data    StructuredContent `json:"data"`
	Message Cursor            `json:"message"`
}
