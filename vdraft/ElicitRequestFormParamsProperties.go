package vdraft

type ElicitRequestFormParamsProperties struct {
	Meta            GetTaskPayloadResult `json:"_meta"`
	Message         Cursor               `json:"message"`
	Mode            Mode                 `json:"mode"`
	RequestedSchema RequestedSchema      `json:"requestedSchema"`
	Task            AnnotationsElement   `json:"task"`
}
