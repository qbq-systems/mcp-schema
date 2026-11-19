package v20260728

type ResultMetaObject struct {
	Description string                     `json:"description"`
	Properties  ResultMetaObjectProperties `json:"properties"`
	Type        TypeElement                `json:"type"`
}
