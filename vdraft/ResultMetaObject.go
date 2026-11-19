package vdraft

type ResultMetaObject struct {
	Description string                     `json:"description"`
	Properties  ResultMetaObjectProperties `json:"properties"`
	Type        TypeElement                `json:"type"`
}
