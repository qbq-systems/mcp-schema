package vdraft

type EmbeddedResourceProperties struct {
	Meta        Meta                `json:"_meta"`
	Annotations AnnotationsClass    `json:"annotations"`
	Resource    *ClientNotification `json:"resource,omitempty"`
	Type        JsonrpcClass        `json:"type"`
	Text        *Cursor             `json:"text,omitempty"`
}
