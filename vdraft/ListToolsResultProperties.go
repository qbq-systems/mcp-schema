package vdraft

type ListToolsResultProperties struct {
	Meta       GetTaskPayloadResult `json:"_meta"`
	NextCursor Cursor               `json:"nextCursor"`
	Tools      Messages             `json:"tools"`
}
