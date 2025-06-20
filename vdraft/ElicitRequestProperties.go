package vdraft

type ElicitRequestProperties struct {
	Method MethodClass  `json:"method"`
	Params IndigoParams `json:"params"`
}
