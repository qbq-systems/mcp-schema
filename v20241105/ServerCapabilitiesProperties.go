package v20241105

type ServerCapabilitiesProperties struct {
	Experimental Experimental `json:"experimental"`
	Logging      Sampling     `json:"logging"`
	Prompts      Roots        `json:"prompts"`
	Resources    Resources    `json:"resources"`
	Tools        Roots        `json:"tools"`
}
