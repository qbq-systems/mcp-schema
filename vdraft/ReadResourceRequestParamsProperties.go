package vdraft

type ReadResourceRequestParamsProperties struct {
	Meta GetTaskPayloadResult `json:"_meta"`
	Task AnnotationsElement   `json:"task"`
	URI  BlobClass            `json:"uri"`
}
