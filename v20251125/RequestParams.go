package v20251125

type RequestParams struct {
	Description string                       `json:"description"`
	Properties  NotificationParamsProperties `json:"properties"`
	Type        AnnotationsType              `json:"type"`
}
