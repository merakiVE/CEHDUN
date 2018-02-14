package parser

type Gateway struct {
	Type string
}

type Pool struct {
	Name string
}

type Event struct {
	Type string
}

type Task struct {
	Name     string
	Type     string
	NeuronID string
	ActionID string
}

type Lane struct {
	Name string
}

type Diagram interface {
	GetGateways() []Gateway
	GetEvents() []Event
	GetTasks() []Task
	GetLanes() []Lane
	GetPools() []Pool

	LoadDiagramByPath(string) error
	LoadDiagramByBuffer([]byte) error
	LoadDiagramByString(string) error
}
