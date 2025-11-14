package vdraft

type LoggingMessageNotificationParamsProperties struct {
	Meta   GetTaskPayloadResult `json:"_meta"`
	Data   PurpleData           `json:"data"`
	Level  AnnotationsElement   `json:"level"`
	Logger Cursor               `json:"logger"`
}
