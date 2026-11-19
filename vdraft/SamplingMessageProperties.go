package vdraft

type SamplingMessageProperties struct {
	Meta    AdditionalProperties `json:"_meta"`
	Content PurpleContent        `json:"content"`
	Role    AdditionalProperties `json:"role"`
}
