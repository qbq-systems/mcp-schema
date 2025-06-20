package vdraft

type ModelPreferencesProperties struct {
	CostPriority         Priority `json:"costPriority"`
	Hints                Audience `json:"hints"`
	IntelligencePriority Priority `json:"intelligencePriority"`
	SpeedPriority        Priority `json:"speedPriority"`
}
