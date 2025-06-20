package vdraft

type Audience struct {
	Description string      `json:"description"`
	Items       EmptyResult `json:"items"`
	Type        string      `json:"type"`
}
