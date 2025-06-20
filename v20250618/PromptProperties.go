package v20250618

type PromptProperties struct {
	Meta        Meta     `json:"_meta"`
	Arguments   Audience `json:"arguments"`
	Description Cursor   `json:"description"`
	Name        Cursor   `json:"name"`
	Title       Cursor   `json:"title"`
}
