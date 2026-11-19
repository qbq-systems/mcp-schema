package v20251125

type ElicitResultProperties struct {
	Meta    GetTaskPayloadResult `json:"_meta"`
	Action  LoggingLevel         `json:"action"`
	Content FluffyContent        `json:"content"`
}
