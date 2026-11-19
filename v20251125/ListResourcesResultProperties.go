package v20251125

type ListResourcesResultProperties struct {
	Meta       GetTaskPayloadResult `json:"_meta"`
	NextCursor Cursor               `json:"nextCursor"`
	Resources  Messages             `json:"resources"`
}
