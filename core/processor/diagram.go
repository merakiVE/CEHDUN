package processor

import (
	"github.com/merakiVE/CVDI/src/models"
	"github.com/merakiVE/CEHDUN/core/parser"
	"github.com/merakiVE/CVDI/core/db"
)

func DiagramProcessor() {}

func VerifyDiagram() {}

func SaveDiagram() ([]map[string]string) {

	diagram := parser.NewProviderParser(parser.PROVIDER_BPMNIO_DIAGRAM)
	diagram.LoadDiagramByPath("/home/hostelix/pruebas_go/diagram.bpmn")

	m := models.ProcedureModel{}

	for index, activity := range diagram.GetActivities() {
		m.Activities = append(m.Activities, models.Activity{
			Name:     activity.Name,
			Type:     activity.Type,
			Sequence: index + 1,
		})
	}

	for _, lane := range diagram.GetLanes() {
		m.Lanes = append(m.Lanes, models.Lane{
			Name: lane.Name,
		})
	}

	m.Owner = "owner_test"

	db.SaveModel(db.GetCurrentDatabase(), &m)

	return m.GetValidationErrors()

}
