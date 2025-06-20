package v20250618

type LoggingMessageNotification struct {
	Description string                               `json:"description"`
	Properties  LoggingMessageNotificationProperties `json:"properties"`
	Required    []string                             `json:"required"`
	Type        AnnotationsType                      `json:"type"`
}
