package vdraft

type AudioContentProperties struct {
	Meta        GetTaskPayloadResult `json:"_meta"`
	Annotations AnnotationsElement   `json:"annotations"`
	Data        BlobClass            `json:"data"`
	MIMEType    Cursor               `json:"mimeType"`
	Type        JsonrpcClass         `json:"type"`
}
