package vdraft

type ResourceProperties struct {
	Meta         AdditionalProperties `json:"_meta"`
	Annotations  ClientResult         `json:"annotations"`
	Description  Cursor               `json:"description"`
	Icons        Audience             `json:"icons"`
	MIMEType     *Cursor              `json:"mimeType,omitempty"`
	Name         Cursor               `json:"name"`
	Size         *Cursor              `json:"size,omitempty"`
	Title        Cursor               `json:"title"`
	URI          *BlobClass           `json:"uri,omitempty"`
	Type         *JsonrpcClass        `json:"type,omitempty"`
	URITemplate  *BlobClass           `json:"uriTemplate,omitempty"`
	InputSchema  *InputSchema         `json:"inputSchema,omitempty"`
	OutputSchema *OutputSchemaClass   `json:"outputSchema,omitempty"`
}
