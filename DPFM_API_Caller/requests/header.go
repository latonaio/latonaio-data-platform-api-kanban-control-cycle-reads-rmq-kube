package requests

type Header struct {
	KanbanControlCycle         string  `json:"KanbanControlCycle"`
	KanbanControlCycleStrategy string  `json:"KanbanControlCycleStrategy"`
	PullRefillInterval         float32 `json:"PullRefillInterval"`
	PullRefillIntervalUnit     string  `json:"PullRefillIntervalUnit"`
	CreationDate               string  `json:"CreationDate"`
	LastChangeDate             string  `json:"LastChangeDate"`
	IsMarkedForDeletion        *bool   `json:"IsMarkedForDeletion"`
}
