package vdraft

type CallToolRequestRequired string

const (
	ID      CallToolRequestRequired = "id"
	Jsonrpc CallToolRequestRequired = "jsonrpc"
	Method  CallToolRequestRequired = "method"
	Params  CallToolRequestRequired = "params"
)
