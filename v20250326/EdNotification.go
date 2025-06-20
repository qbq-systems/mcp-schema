package v20250326

type EdNotification struct {
	Description string                            `json:"description"`
	Properties  InitializedNotificationProperties `json:"properties"`
	Required    []string                          `json:"required"`
	Type        AnnotationsType                   `json:"type"`
}
