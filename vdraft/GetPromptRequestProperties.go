package vdraft

type GetPromptRequestProperties struct {
	ID      EmptyResult    `json:"id"`
	Jsonrpc JsonrpcClass   `json:"jsonrpc"`
	Method  JsonrpcClass   `json:"method"`
	Params  IndecentParams `json:"params"`
}
