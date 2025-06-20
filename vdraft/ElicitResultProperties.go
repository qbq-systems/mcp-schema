package vdraft

type ElicitResultProperties struct {
	Meta    Meta         `json:"_meta"`
	Action  LoggingLevel `json:"action"`
	Content ContentClass `json:"content"`
}
