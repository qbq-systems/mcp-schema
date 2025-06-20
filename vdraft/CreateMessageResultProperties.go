package vdraft

type CreateMessageResultProperties struct {
	Meta       Meta               `json:"_meta"`
	Content    ClientNotification `json:"content"`
	Model      Cursor             `json:"model"`
	Role       EmptyResult        `json:"role"`
	StopReason Cursor             `json:"stopReason"`
}
