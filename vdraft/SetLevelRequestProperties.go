package vdraft

type SetLevelRequestProperties struct {
	ID      EmptyResult       `json:"id"`
	Jsonrpc JsonrpcClass      `json:"jsonrpc"`
	Method  JsonrpcClass      `json:"method"`
	Params  MischievousParams `json:"params"`
}
