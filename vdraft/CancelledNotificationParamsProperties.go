package vdraft

type CancelledNotificationParamsProperties struct {
	Meta      Meta             `json:"_meta"`
	Reason    Cursor           `json:"reason"`
	RequestID AnnotationsClass `json:"requestId"`
}
