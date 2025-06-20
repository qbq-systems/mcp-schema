package v20250326

type Notification struct {
	Properties NotificationProperties `json:"properties"`
	Required   []string               `json:"required"`
	Type       AnnotationsType        `json:"type"`
}
