package v20250326

type ResourceProperties struct {
	Annotations AnnotationsClass `json:"annotations"`
	Description Cursor           `json:"description"`
	MIMEType    Cursor           `json:"mimeType"`
	Name        Cursor           `json:"name"`
	Size        Cursor           `json:"size"`
	URI         BlobClass        `json:"uri"`
}
