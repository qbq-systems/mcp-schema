package v20241105

type ResourceProperties struct {
	Annotations Annotations `json:"annotations"`
	Description Cursor      `json:"description"`
	MIMEType    Cursor      `json:"mimeType"`
	Name        Cursor      `json:"name"`
	Size        Cursor      `json:"size"`
	URI         Blob        `json:"uri"`
}
