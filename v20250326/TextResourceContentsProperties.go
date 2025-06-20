package v20250326

type TextResourceContentsProperties struct {
	MIMEType Cursor    `json:"mimeType"`
	Text     Cursor    `json:"text"`
	URI      BlobClass `json:"uri"`
}
