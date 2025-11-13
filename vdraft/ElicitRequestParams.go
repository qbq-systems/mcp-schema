package vdraft

type ElicitRequestParams struct {
	AnyOf       []EmptyResult `json:"anyOf"`
	Description string        `json:"description"`
}
