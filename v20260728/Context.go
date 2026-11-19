package v20260728

type Context struct {
	Description string            `json:"description"`
	Properties  ContextProperties `json:"properties"`
	Type        TypeElement       `json:"type"`
}
