package v20260728

type ServerCapabilitiesProperties struct {
	Completions  ClientResult  `json:"completions"`
	Experimental InputRequests `json:"experimental"`
	Extensions   InputRequests `json:"extensions"`
	Logging      ClientResult  `json:"logging"`
	Prompts      Prompts       `json:"prompts"`
	Resources    Resources     `json:"resources"`
	Tools        Prompts       `json:"tools"`
}
