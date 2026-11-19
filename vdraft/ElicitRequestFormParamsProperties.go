package vdraft

type ElicitRequestFormParamsProperties struct {
	Meta            AdditionalProperties `json:"_meta"`
	Message         Cursor               `json:"message"`
	Mode            Mode                 `json:"mode"`
	RequestedSchema *RequestedSchema     `json:"requestedSchema,omitempty"`
	URL             *BlobClass           `json:"url,omitempty"`
}
