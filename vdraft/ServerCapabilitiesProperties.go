package vdraft

type ServerCapabilitiesProperties struct {
	Completions  Metadata     `json:"completions"`
	Experimental Experimental `json:"experimental"`
	Logging      Metadata     `json:"logging"`
	Prompts      Roots        `json:"prompts"`
	Resources    Resources    `json:"resources"`
	Tools        Roots        `json:"tools"`
}
