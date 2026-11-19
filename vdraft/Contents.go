package vdraft

type Contents struct {
	Items ClientRequest `json:"items"`
	Type  AudienceType  `json:"type"`
}
