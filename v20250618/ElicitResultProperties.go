package v20250618

type ElicitResultProperties struct {
	Meta    Meta         `json:"_meta"`
	Action  LoggingLevel `json:"action"`
	Content ContentClass `json:"content"`
}
