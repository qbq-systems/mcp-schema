package v20251125

type ListResourceTemplatesResultProperties struct {
	Meta              GetTaskPayloadResult `json:"_meta"`
	NextCursor        Cursor               `json:"nextCursor"`
	ResourceTemplates Messages             `json:"resourceTemplates"`
}
