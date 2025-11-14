package vdraft

type ElicitResultProperties struct {
	Meta    GetTaskPayloadResult `json:"_meta"`
	Action  LoggingLevel         `json:"action"`
	Content FluffyContent        `json:"content"`
}
