package v20250326

type CreateMessageResultProperties struct {
	Meta       PurpleMeta         `json:"_meta"`
	Content    ClientNotification `json:"content"`
	Model      Cursor             `json:"model"`
	Role       EmptyResult        `json:"role"`
	StopReason Cursor             `json:"stopReason"`
}
