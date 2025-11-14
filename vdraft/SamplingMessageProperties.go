package vdraft

type SamplingMessageProperties struct {
	Meta    GetTaskPayloadResult `json:"_meta"`
	Content PurpleContent        `json:"content"`
	Role    EmptyResult          `json:"role"`
}
