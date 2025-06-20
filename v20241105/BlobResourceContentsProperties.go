package v20241105

type BlobResourceContentsProperties struct {
	Blob     Blob   `json:"blob"`
	MIMEType Cursor `json:"mimeType"`
	URI      Blob   `json:"uri"`
}
