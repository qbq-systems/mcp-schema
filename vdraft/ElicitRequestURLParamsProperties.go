package vdraft

type ElicitRequestURLParamsProperties struct {
	Meta          GetTaskPayloadResult `json:"_meta"`
	ElicitationID Cursor               `json:"elicitationId"`
	Message       Cursor               `json:"message"`
	Mode          Mode                 `json:"mode"`
	Task          AnnotationsElement   `json:"task"`
	URL           BlobClass            `json:"url"`
}
