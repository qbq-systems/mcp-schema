package v20260728

type RequestMetaObject struct {
	Description string                      `json:"description"`
	Properties  RequestMetaObjectProperties `json:"properties"`
	Required    []string                    `json:"required"`
	Type        TypeElement                 `json:"type"`
}
