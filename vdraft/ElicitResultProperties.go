package vdraft

type ElicitResultProperties struct {
	Meta    AdditionalProperties `json:"_meta"`
	Action  LoggingLevel         `json:"action"`
	Content FluffyContent        `json:"content"`
}
