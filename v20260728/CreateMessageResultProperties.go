package v20260728

type CreateMessageResultProperties struct {
	Meta       AdditionalProperties `json:"_meta"`
	Content    PurpleContent        `json:"content"`
	Model      Cursor               `json:"model"`
	Role       AdditionalProperties `json:"role"`
	StopReason Cursor               `json:"stopReason"`
}
