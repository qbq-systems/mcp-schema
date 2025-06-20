package v20241105

type CallToolResultProperties struct {
	Meta    PurpleMeta `json:"_meta"`
	Content Content    `json:"content"`
	IsError Cursor     `json:"isError"`
}
