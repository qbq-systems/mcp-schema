package v20250326

type BlobResourceContentsProperties struct {
	Blob     BlobClass `json:"blob"`
	MIMEType Cursor    `json:"mimeType"`
	URI      BlobClass `json:"uri"`
}
