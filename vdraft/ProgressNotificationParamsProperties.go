package vdraft

type ProgressNotificationParamsProperties struct {
	Meta          GetTaskPayloadResult `json:"_meta"`
	Message       Cursor               `json:"message"`
	Progress      Cursor               `json:"progress"`
	ProgressToken AnnotationsElement   `json:"progressToken"`
	Total         Cursor               `json:"total"`
}
