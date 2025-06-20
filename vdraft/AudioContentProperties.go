package vdraft

type AudioContentProperties struct {
	Meta        Meta             `json:"_meta"`
	Annotations AnnotationsClass `json:"annotations"`
	Data        BlobClass        `json:"data"`
	MIMEType    Cursor           `json:"mimeType"`
	Type        MethodClass      `json:"type"`
}
