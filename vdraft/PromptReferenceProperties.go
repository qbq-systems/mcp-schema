package vdraft

type PromptReferenceProperties struct {
	Name  Cursor       `json:"name"`
	Title Cursor       `json:"title"`
	Type  JsonrpcClass `json:"type"`
}
