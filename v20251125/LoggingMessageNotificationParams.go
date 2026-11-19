package v20251125

type LoggingMessageNotificationParams struct {
	Description string                                     `json:"description"`
	Properties  LoggingMessageNotificationParamsProperties `json:"properties"`
	Required    []string                                   `json:"required"`
	Type        AnnotationsType                            `json:"type"`
}
