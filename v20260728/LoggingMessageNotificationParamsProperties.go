package v20260728

type LoggingMessageNotificationParamsProperties struct {
	Meta   AdditionalProperties `json:"_meta"`
	Data   StructuredContent    `json:"data"`
	Level  ClientResult         `json:"level"`
	Logger Cursor               `json:"logger"`
}
