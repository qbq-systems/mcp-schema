package v20260728

type BlobResourceContentsProperties struct {
	Meta     AdditionalProperties `json:"_meta"`
	Blob     *BlobClass           `json:"blob,omitempty"`
	MIMEType Cursor               `json:"mimeType"`
	URI      BlobClass            `json:"uri"`
	Text     *Cursor              `json:"text,omitempty"`
}
