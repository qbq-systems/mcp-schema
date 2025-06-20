package vdraft

type ListRootsResultProperties struct {
	Meta  Meta     `json:"_meta"`
	Roots Messages `json:"roots"`
}
