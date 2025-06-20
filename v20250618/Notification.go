package v20250618

type Notification struct {
	Properties NotificationProperties `json:"properties"`
	Required   []string               `json:"required"`
	Type       AnnotationsType        `json:"type"`
}
