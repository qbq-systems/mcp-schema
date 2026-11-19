package vdraft

type ProgressNotificationParamsProperties struct {
	Meta          AdditionalProperties `json:"_meta"`
	Message       Cursor               `json:"message"`
	Progress      Cursor               `json:"progress"`
	ProgressToken ClientResult         `json:"progressToken"`
	Total         Cursor               `json:"total"`
}
