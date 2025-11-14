package vdraft

type ListTasksResultProperties struct {
	Meta       GetTaskPayloadResult `json:"_meta"`
	NextCursor Cursor               `json:"nextCursor"`
	Tasks      Messages             `json:"tasks"`
}
