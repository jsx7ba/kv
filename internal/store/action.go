package store

// Action defines which synchronization actions might happen between nodes in the cluster.
type Action int

const (
	Put Action = iota + 1
	Delete
)

// Entry Defines
type Entry struct {
	Action Action      `json:"action"`
	Key    string      `json:"key"`
	Value  interface{} `json:"value,omitempty"`
}
