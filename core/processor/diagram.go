package processor

import "github.com/merakiVE/CVDI/src/models"
import (
	"github.com/merakiVE/CEHDUN/core/parser"
	"github.com/merakiVE/CVDI/core/db"
)

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

func DiagramProcessor() {

}

func VerifyDiagram() {}

func SaveDiagram() ([]map[string]string) {
	diagram := parser.NewDiagram()
	diagram.ReadFromFile("/home/hostelix/pruebas_go/diagram.bpmn")

	m := models.ProcedureModel{}

	for index, activity := range diagram.GetSuccessionProcess() {

		m.Activities = append(m.Activities, models.Activity{
			Name:     diagram.GetAttributeElement(activity, parser.ATTR_NAME),
			Sequence: index + 1,
		})
	}

	for _, lane := range diagram.GetLanes() {
		m.Lanes = append(m.Lanes, models.Lane{
			Name: diagram.GetAttributeElement(lane, parser.ATTR_NAME),
		})
	}

	m.Owner = "owner_test"

	db.SaveModel(db.GetCurrentDatabase(), &m)


	return m.GetValidationErrors()

}
