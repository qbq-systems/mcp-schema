package vdraft

type CreateMessageResultProperties struct {
	Meta       GetTaskPayloadResult `json:"_meta"`
	Content    PurpleContent        `json:"content"`
	Model      Cursor               `json:"model"`
	Role       EmptyResult          `json:"role"`
	StopReason Cursor               `json:"stopReason"`
}
