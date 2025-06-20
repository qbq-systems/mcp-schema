package v20241105

type TextResourceContentsProperties struct {
	MIMEType Cursor `json:"mimeType"`
	Text     Cursor `json:"text"`
	URI      Blob   `json:"uri"`
}
