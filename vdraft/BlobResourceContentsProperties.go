package vdraft

type BlobResourceContentsProperties struct {
	Meta     GetTaskPayloadResult `json:"_meta"`
	Blob     *BlobClass           `json:"blob,omitempty"`
	MIMEType Cursor               `json:"mimeType"`
	URI      BlobClass            `json:"uri"`
	Text     *Cursor              `json:"text,omitempty"`
}
