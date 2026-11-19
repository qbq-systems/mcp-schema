package v20251125

type CancelledNotificationParamsProperties struct {
	Meta      GetTaskPayloadResult `json:"_meta"`
	Reason    Cursor               `json:"reason"`
	RequestID AnnotationsElement   `json:"requestId"`
}
