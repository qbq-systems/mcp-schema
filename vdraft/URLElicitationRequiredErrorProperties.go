package vdraft

type URLElicitationRequiredErrorProperties struct {
	Error   ErrorClass   `json:"error"`
	ID      EmptyResult  `json:"id"`
	Jsonrpc JsonrpcClass `json:"jsonrpc"`
}
