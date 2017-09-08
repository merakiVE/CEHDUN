package parser

/* Constantes en donde se definen los proveedores de parseo */
const (
	PROVIDER_BPMNIO_DIAGRAM = "bpmn.io"
	PROVIDER_GOJS_DIAGRAM   = "gojs"
	PROVIDER_DRAWIO_DIAGRAM = "draw.io"
)

/*
	Funcion que sirve para instanciar el nuevo parseador dependiendo del proveedor que se especifique

	Retorna una interface{}, los metodos de esta deberian coincidir con la interface Diagram
 */
func NewProviderParser(provider string) (Diagram) {

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
