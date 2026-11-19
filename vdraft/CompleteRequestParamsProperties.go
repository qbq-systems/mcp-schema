package vdraft

type CompleteRequestParamsProperties struct {
	Meta     AdditionalProperties `json:"_meta"`
	Argument Argument             `json:"argument"`
	Context  Context              `json:"context"`
	Ref      ClientRequest        `json:"ref"`
}
