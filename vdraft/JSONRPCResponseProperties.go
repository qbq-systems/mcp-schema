package vdraft

type JSONRPCResponseProperties struct {
	ID      EmptyResult `json:"id"`
	Jsonrpc MethodClass `json:"jsonrpc"`
	Result  EmptyResult `json:"result"`
}
