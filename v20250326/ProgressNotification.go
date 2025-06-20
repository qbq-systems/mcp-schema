package v20250326

type ProgressNotification struct {
	Description string                         `json:"description"`
	Properties  ProgressNotificationProperties `json:"properties"`
	Required    []string                       `json:"required"`
	Type        AnnotationsType                `json:"type"`
}
