package vdraft

type ToolExecution struct {
	Description string                  `json:"description"`
	Properties  ToolExecutionProperties `json:"properties"`
	Type        AnnotationsType         `json:"type"`
}
