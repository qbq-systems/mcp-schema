package vdraft

type CompleteResultProperties struct {
	Meta       GetTaskPayloadResult `json:"_meta"`
	Completion Completion           `json:"completion"`
}
