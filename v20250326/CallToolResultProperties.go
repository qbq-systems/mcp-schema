package v20250326

type CallToolResultProperties struct {
	Meta    PurpleMeta   `json:"_meta"`
	Content ContentClass `json:"content"`
	IsError Cursor       `json:"isError"`
}
