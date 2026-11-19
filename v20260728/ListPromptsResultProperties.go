package v20260728

type ListPromptsResultProperties struct {
	Meta              AdditionalProperties `json:"_meta"`
	CacheScope        LoggingLevel         `json:"cacheScope"`
	NextCursor        Cursor               `json:"nextCursor"`
	Prompts           *JSONArray           `json:"prompts,omitempty"`
	ResultType        Cursor               `json:"resultType"`
	TTLMS             Priority             `json:"ttlMs"`
	ResourceTemplates *JSONArray           `json:"resourceTemplates,omitempty"`
	Resources         *JSONArray           `json:"resources,omitempty"`
	Tools             *JSONArray           `json:"tools,omitempty"`
}
