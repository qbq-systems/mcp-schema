package vdraft

type JsonrpcMessage struct {
	AnyOf       []EmptyResult `json:"anyOf"`
	Description string        `json:"description"`
}
