package vdraft

type Mode struct {
	Const       string      `json:"const"`
	Description string      `json:"description"`
	Type        TypeElement `json:"type"`
}
