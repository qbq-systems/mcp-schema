package v20250618

type ServerCapabilitiesProperties struct {
	Completions  Elicitation  `json:"completions"`
	Experimental Experimental `json:"experimental"`
	Logging      Elicitation  `json:"logging"`
	Prompts      Roots        `json:"prompts"`
	Resources    Resources    `json:"resources"`
	Tools        Roots        `json:"tools"`
}
