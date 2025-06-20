package v20250618

type EmbeddedResourceProperties struct {
	Meta        Meta                `json:"_meta"`
	Annotations AnnotationsClass    `json:"annotations"`
	Resource    *ClientNotification `json:"resource,omitempty"`
	Type        MethodClass         `json:"type"`
	Text        *Cursor             `json:"text,omitempty"`
}
