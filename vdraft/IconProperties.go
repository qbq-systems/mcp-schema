package vdraft

type IconProperties struct {
	MIMEType Cursor            `json:"mimeType"`
	Sizes    SupportedVersions `json:"sizes"`
	Src      BlobClass         `json:"src"`
	Theme    LoggingLevel      `json:"theme"`
}
