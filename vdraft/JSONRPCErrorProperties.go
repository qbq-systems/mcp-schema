package vdraft

type JSONRPCErrorProperties struct {
	Error   Error       `json:"error"`
	ID      EmptyResult `json:"id"`
	Jsonrpc MethodClass `json:"jsonrpc"`
}
