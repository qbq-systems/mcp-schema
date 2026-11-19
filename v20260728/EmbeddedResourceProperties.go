package v20260728

type EmbeddedResourceProperties struct {
	Meta        AdditionalProperties `json:"_meta"`
	Annotations ClientResult         `json:"annotations"`
	Resource    *ClientRequest       `json:"resource,omitempty"`
	Type        JsonrpcClass         `json:"type"`
	Text        *Cursor              `json:"text,omitempty"`
}
