package vdraft

type Contents struct {
	Items ClientNotification `json:"items"`
	Type  string             `json:"type"`
}
