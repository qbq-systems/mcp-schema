package vdraft

type RequestParams struct {
	Description string                  `json:"description"`
	Properties  RequestParamsProperties `json:"properties"`
	Type        AnnotationsType         `json:"type"`
}
