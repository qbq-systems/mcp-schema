package vdraft

type ClientCapabilitiesProperties struct {
	Elicitation  Elicitation   `json:"elicitation"`
	Experimental InputRequests `json:"experimental"`
	Extensions   InputRequests `json:"extensions"`
	Roots        Roots         `json:"roots"`
	Sampling     Sampling      `json:"sampling"`
}
