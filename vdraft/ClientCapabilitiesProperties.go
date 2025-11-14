package vdraft

type ClientCapabilitiesProperties struct {
	Elicitation  PurpleElicitation `json:"elicitation"`
	Experimental Experimental      `json:"experimental"`
	Roots        Roots             `json:"roots"`
	Sampling     PurpleSampling    `json:"sampling"`
	Tasks        PurpleTasks       `json:"tasks"`
}
