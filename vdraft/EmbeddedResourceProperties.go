package vdraft

type EmbeddedResourceProperties struct {
	Meta        GetTaskPayloadResult `json:"_meta"`
	Annotations AnnotationsElement   `json:"annotations"`
	Resource    *ClientNotification  `json:"resource,omitempty"`
	Type        JsonrpcClass         `json:"type"`
	Text        *Cursor              `json:"text,omitempty"`
}
