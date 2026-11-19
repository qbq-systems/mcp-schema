package v20260728

type NotificationMetaObject struct {
	Description string                           `json:"description"`
	Properties  NotificationMetaObjectProperties `json:"properties"`
	Type        TypeElement                      `json:"type"`
}
