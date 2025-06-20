package v20241105

type ImageContentProperties struct {
	Annotations Annotations `json:"annotations"`
	Data        Blob        `json:"data"`
	MIMEType    Cursor      `json:"mimeType"`
	Type        Method      `json:"type"`
}
