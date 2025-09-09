package vdraft

type ResourceTemplateProperties struct {
	Meta        Meta             `json:"_meta"`
	Annotations AnnotationsClass `json:"annotations"`
	Description Cursor           `json:"description"`
	MIMEType    Cursor           `json:"mimeType"`
	Name        Cursor           `json:"name"`
	Title       Cursor           `json:"title"`
	URITemplate BlobClass        `json:"uriTemplate"`
}
