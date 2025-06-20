package vdraft

type Priority struct {
	Description string      `json:"description"`
	Maximum     int64       `json:"maximum"`
	Minimum     int64       `json:"minimum"`
	Type        TypeElement `json:"type"`
}
