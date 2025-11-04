package vdraft

type ProgressNotificationParamsProperties struct {
	Meta          Meta             `json:"_meta"`
	Message       Cursor           `json:"message"`
	Progress      Cursor           `json:"progress"`
	ProgressToken AnnotationsClass `json:"progressToken"`
	Total         Cursor           `json:"total"`
}
