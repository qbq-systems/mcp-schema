package v20250326

type ModelPreferencesProperties struct {
	CostPriority         Priority `json:"costPriority"`
	Hints                Audience `json:"hints"`
	IntelligencePriority Priority `json:"intelligencePriority"`
	SpeedPriority        Priority `json:"speedPriority"`
}
