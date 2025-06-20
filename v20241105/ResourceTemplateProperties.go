package v20241105

type ResourceTemplateProperties struct {
	Annotations Annotations `json:"annotations"`
	Description Cursor      `json:"description"`
	MIMEType    Cursor      `json:"mimeType"`
	Name        Cursor      `json:"name"`
	URITemplate Blob        `json:"uriTemplate"`
}
