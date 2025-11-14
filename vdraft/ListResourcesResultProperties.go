package vdraft

type ListResourcesResultProperties struct {
	Meta       GetTaskPayloadResult `json:"_meta"`
	NextCursor Cursor               `json:"nextCursor"`
	Resources  Messages             `json:"resources"`
}
