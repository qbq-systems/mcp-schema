package v20250618

type BlobResourceContentsProperties struct {
	Meta     Meta       `json:"_meta"`
	Blob     *BlobClass `json:"blob,omitempty"`
	MIMEType Cursor     `json:"mimeType"`
	URI      BlobClass  `json:"uri"`
	Text     *Cursor    `json:"text,omitempty"`
}
