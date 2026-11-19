package v20251125

type IconProperties struct {
	MIMEType Cursor       `json:"mimeType"`
	Sizes    Sizes        `json:"sizes"`
	Src      BlobClass    `json:"src"`
	Theme    LoggingLevel `json:"theme"`
}
