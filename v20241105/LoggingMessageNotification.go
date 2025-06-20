package v20241105

type LoggingMessageNotification struct {
	Description string                               `json:"description"`
	Properties  LoggingMessageNotificationProperties `json:"properties"`
	Required    []string                             `json:"required"`
	Type        AnnotationsType                      `json:"type"`
}
