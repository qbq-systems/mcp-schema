package vdraft

type Values struct {
	Description string  `json:"description"`
	Items       Default `json:"items"`
	Type        string  `json:"type"`
}
