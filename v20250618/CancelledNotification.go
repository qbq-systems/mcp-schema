package v20250618

type CancelledNotification struct {
	Description string                          `json:"description"`
	Properties  CancelledNotificationProperties `json:"properties"`
	Required    []string                        `json:"required"`
	Type        AnnotationsType                 `json:"type"`
}
