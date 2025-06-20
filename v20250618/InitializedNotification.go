package v20250618

type InitializedNotification struct {
	Description string                            `json:"description"`
	Properties  InitializedNotificationProperties `json:"properties"`
	Required    []string                          `json:"required"`
	Type        AnnotationsType                   `json:"type"`
}
