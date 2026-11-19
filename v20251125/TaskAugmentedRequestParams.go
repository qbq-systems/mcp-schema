package v20251125

type TaskAugmentedRequestParams struct {
	Description string                               `json:"description"`
	Properties  TaskAugmentedRequestParamsProperties `json:"properties"`
	Type        AnnotationsType                      `json:"type"`
}
