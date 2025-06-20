package vdraft

type AnnotationsProperties struct {
	Audience     Audience `json:"audience"`
	LastModified Cursor   `json:"lastModified"`
	Priority     Priority `json:"priority"`
}
