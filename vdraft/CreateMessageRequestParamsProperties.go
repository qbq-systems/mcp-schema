package vdraft

type CreateMessageRequestParamsProperties struct {
	Meta             GetTaskPayloadResult `json:"_meta"`
	IncludeContext   LoggingLevel         `json:"includeContext"`
	MaxTokens        Cursor               `json:"maxTokens"`
	Messages         Messages             `json:"messages"`
	Metadata         Metadata             `json:"metadata"`
	ModelPreferences AnnotationsElement   `json:"modelPreferences"`
	StopSequences    StopSequences        `json:"stopSequences"`
	SystemPrompt     Cursor               `json:"systemPrompt"`
	Task             AnnotationsElement   `json:"task"`
	Temperature      Default              `json:"temperature"`
	ToolChoice       AnnotationsElement   `json:"toolChoice"`
	Tools            Audience             `json:"tools"`
}
