package vdraft

type ProgressToken struct {
	Description string        `json:"description"`
	Type        []TypeElement `json:"type"`
}
