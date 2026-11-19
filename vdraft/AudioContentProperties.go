package vdraft

type AudioContentProperties struct {
	Meta        AdditionalProperties `json:"_meta"`
	Annotations ClientResult         `json:"annotations"`
	Data        BlobClass            `json:"data"`
	MIMEType    Cursor               `json:"mimeType"`
	Type        JsonrpcClass         `json:"type"`
}
