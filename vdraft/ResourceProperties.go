package vdraft

type ResourceProperties struct {
	Meta        Meta             `json:"_meta"`
	Annotations AnnotationsClass `json:"annotations"`
	Description Cursor           `json:"description"`
	MIMEType    Cursor           `json:"mimeType"`
	Name        Cursor           `json:"name"`
	Size        *Cursor          `json:"size,omitempty"`
	Title       Cursor           `json:"title"`
	URI         *BlobClass       `json:"uri,omitempty"`
	Type        *JsonrpcClass    `json:"type,omitempty"`
	URITemplate *BlobClass       `json:"uriTemplate,omitempty"`
}
