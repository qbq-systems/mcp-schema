package vdraft

type LoggingMessageNotificationParamsProperties struct {
	Meta   AdditionalProperties `json:"_meta"`
	Data   StructuredContent    `json:"data"`
	Level  ClientResult         `json:"level"`
	Logger Cursor               `json:"logger"`
}
