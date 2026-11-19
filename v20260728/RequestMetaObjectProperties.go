package v20260728

type RequestMetaObjectProperties struct {
	IoModelcontextprotocolClientCapabilities ClientResult `json:"io.modelcontextprotocol/clientCapabilities"`
	IoModelcontextprotocolClientInfo         ClientResult `json:"io.modelcontextprotocol/clientInfo"`
	IoModelcontextprotocolLogLevel           ClientResult `json:"io.modelcontextprotocol/logLevel"`
	IoModelcontextprotocolProtocolVersion    Cursor       `json:"io.modelcontextprotocol/protocolVersion"`
	ProgressToken                            ClientResult `json:"progressToken"`
}
