package v20260728

type CreateMessageRequestParamsProperties struct {
	IncludeContext   LoggingLevel  `json:"includeContext"`
	MaxTokens        Cursor        `json:"maxTokens"`
	Messages         JSONArray     `json:"messages"`
	Metadata         ClientResult  `json:"metadata"`
	ModelPreferences ClientResult  `json:"modelPreferences"`
	StopSequences    StopSequences `json:"stopSequences"`
	SystemPrompt     Cursor        `json:"systemPrompt"`
	Temperature      Default       `json:"temperature"`
	ToolChoice       ClientResult  `json:"toolChoice"`
	Tools            Audience      `json:"tools"`
}
