package parser

type Gateway struct {
	Type string
}

type Event struct {
	Type string
}

type Activity struct {
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
	GetActivities() []Activity
	GetLanes() []Lane
	//GetPools()

	LoadDiagramByPath(string)
	LoadDiagramByBuffer([]byte)
	LoadDiagramByString(string)
}
