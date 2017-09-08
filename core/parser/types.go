package parser

type Gateway struct {
	Type string
}

type Event struct {
	Type string
}

type Activity struct {
	Name string
}

type Diagram interface {
	GetGateways() []Gateway
	GetEvents() []Event
	GetActivities() []Activity
	//GetLanes()
	//GetPools()
}
