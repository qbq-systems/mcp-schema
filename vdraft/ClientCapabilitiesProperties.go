package vdraft

type ClientCapabilitiesProperties struct {
	Elicitation  Elicitation  `json:"elicitation"`
	Experimental Experimental `json:"experimental"`
	Roots        Roots        `json:"roots"`
	Sampling     Elicitation  `json:"sampling"`
}
