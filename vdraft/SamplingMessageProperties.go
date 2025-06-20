package vdraft

type SamplingMessageProperties struct {
	Content ClientNotification `json:"content"`
	Role    EmptyResult        `json:"role"`
}
