package vdraft

type ListResourcesResultProperties struct {
	Meta       Meta     `json:"_meta"`
	NextCursor Cursor   `json:"nextCursor"`
	Resources  Messages `json:"resources"`
}
