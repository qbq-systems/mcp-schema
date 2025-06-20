package v20241105

type ModelPreferencesProperties struct {
	CostPriority         Priority `json:"costPriority"`
	Hints                Hints    `json:"hints"`
	IntelligencePriority Priority `json:"intelligencePriority"`
	SpeedPriority        Priority `json:"speedPriority"`
}
