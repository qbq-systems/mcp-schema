package v20250326

type AudioContentProperties struct {
	Annotations AnnotationsClass `json:"annotations"`
	Data        BlobClass        `json:"data"`
	MIMEType    Cursor           `json:"mimeType"`
	Type        TypeClass        `json:"type"`
}
