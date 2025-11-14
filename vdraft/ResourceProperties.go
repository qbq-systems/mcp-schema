package vdraft

type ResourceProperties struct {
	Meta         GetTaskPayloadResult `json:"_meta"`
	Annotations  AnnotationsElement   `json:"annotations"`
	Description  Cursor               `json:"description"`
	Icons        Audience             `json:"icons"`
	MIMEType     *Cursor              `json:"mimeType,omitempty"`
	Name         Cursor               `json:"name"`
	Size         *Cursor              `json:"size,omitempty"`
	Title        Cursor               `json:"title"`
	URI          *BlobClass           `json:"uri,omitempty"`
	Type         *JsonrpcClass        `json:"type,omitempty"`
	URITemplate  *BlobClass           `json:"uriTemplate,omitempty"`
	InputSchema  *PutSchema           `json:"inputSchema,omitempty"`
	OutputSchema *PutSchema           `json:"outputSchema,omitempty"`
}
