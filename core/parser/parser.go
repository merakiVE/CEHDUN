package parser

const (
	PROVIDER_BPMNIO_DIAGRAM = "bpmn.io"
	PROVIDER_GOJS_DIAGRAM   = "gojs"
	PROVIDER_DRAWIO_DIAGRAM = "draw.io"
)

func NewProviderParser(provider string) (interface{}) {

	switch provider {

	case PROVIDER_BPMNIO_DIAGRAM:
		return NewParserBPMNIO()

	case PROVIDER_GOJS_DIAGRAM:
		return nil

	case PROVIDER_DRAWIO_DIAGRAM:
		return nil

	default:
		return NewParserBPMNIO()
	}
	return nil
}
