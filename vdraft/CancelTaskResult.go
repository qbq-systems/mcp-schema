package vdraft

type CancelTaskResult struct {
	AllOf       []EmptyResult `json:"allOf"`
	Description string        `json:"description"`
}
