package vdraft

type ReadResourceResultProperties struct {
	Meta     GetTaskPayloadResult `json:"_meta"`
	Contents Contents             `json:"contents"`
}
