package vdraft

type PromptProperties struct {
	Meta        Meta     `json:"_meta"`
	Arguments   Audience `json:"arguments"`
	Description Cursor   `json:"description"`
	Icons       Audience `json:"icons"`
	Name        Cursor   `json:"name"`
	Title       Cursor   `json:"title"`
}
