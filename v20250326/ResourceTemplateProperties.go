package v20250326

type ResourceTemplateProperties struct {
	Annotations AnnotationsClass `json:"annotations"`
	Description Cursor           `json:"description"`
	MIMEType    Cursor           `json:"mimeType"`
	Name        Cursor           `json:"name"`
	URITemplate BlobClass        `json:"uriTemplate"`
}
