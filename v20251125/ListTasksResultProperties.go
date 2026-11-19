package v20251125

type ListTasksResultProperties struct {
	Meta       GetTaskPayloadResult `json:"_meta"`
	NextCursor Cursor               `json:"nextCursor"`
	Tasks      Messages             `json:"tasks"`
}
