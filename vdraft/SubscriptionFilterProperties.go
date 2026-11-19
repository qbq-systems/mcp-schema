package vdraft

type SubscriptionFilterProperties struct {
	PromptsListChanged    Cursor            `json:"promptsListChanged"`
	ResourceSubscriptions SupportedVersions `json:"resourceSubscriptions"`
	ResourcesListChanged  Cursor            `json:"resourcesListChanged"`
	ToolsListChanged      Cursor            `json:"toolsListChanged"`
}
