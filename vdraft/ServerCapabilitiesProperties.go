package vdraft

type ServerCapabilitiesProperties struct {
	Completions  Sampling     `json:"completions"`
	Experimental Experimental `json:"experimental"`
	Logging      Sampling     `json:"logging"`
	Prompts      Roots        `json:"prompts"`
	Resources    Resources    `json:"resources"`
	Tools        Roots        `json:"tools"`
}
