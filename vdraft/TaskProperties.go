package vdraft

type TaskProperties struct {
	CreatedAt     Cursor             `json:"createdAt"`
	LastUpdatedAt Cursor             `json:"lastUpdatedAt"`
	PollInterval  Cursor             `json:"pollInterval"`
	Status        AnnotationsElement `json:"status"`
	StatusMessage Cursor             `json:"statusMessage"`
	TaskID        Cursor             `json:"taskId"`
	TTL           Cursor             `json:"ttl"`
}
