package v20241105

type ProgressNotification struct {
	Description string                         `json:"description"`
	Properties  ProgressNotificationProperties `json:"properties"`
	Required    []string                       `json:"required"`
	Type        AnnotationsType                `json:"type"`
}
