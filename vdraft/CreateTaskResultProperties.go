package vdraft

type CreateTaskResultProperties struct {
	Meta GetTaskPayloadResult `json:"_meta"`
	Task EmptyResult          `json:"task"`
}
