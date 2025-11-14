package vdraft

type ListRootsResultProperties struct {
	Meta  GetTaskPayloadResult `json:"_meta"`
	Roots Messages             `json:"roots"`
}
