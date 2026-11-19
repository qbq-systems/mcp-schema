package v20251125

type ListToolsResultProperties struct {
	Meta       GetTaskPayloadResult `json:"_meta"`
	NextCursor Cursor               `json:"nextCursor"`
	Tools      Messages             `json:"tools"`
}
