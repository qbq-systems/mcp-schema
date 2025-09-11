package vdraft

type Contents struct {
	Items ClientNotification `json:"items"`
	Type  AudienceType       `json:"type"`
}
