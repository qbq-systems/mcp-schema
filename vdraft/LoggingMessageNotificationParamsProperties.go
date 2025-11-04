package vdraft

type LoggingMessageNotificationParamsProperties struct {
	Meta   Meta             `json:"_meta"`
	Data   PurpleData       `json:"data"`
	Level  AnnotationsClass `json:"level"`
	Logger Cursor           `json:"logger"`
}
