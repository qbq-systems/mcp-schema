package vdraft

type NotificationMetaObject struct {
	Description string                           `json:"description"`
	Properties  NotificationMetaObjectProperties `json:"properties"`
	Type        TypeElement                      `json:"type"`
}
