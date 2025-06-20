package vdraft

type Notification struct {
	Properties NotificationProperties `json:"properties"`
	Required   []string               `json:"required"`
	Type       AnnotationsType        `json:"type"`
}
