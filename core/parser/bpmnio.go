package parser

import (
	"strings"
	"github.com/beevik/etree"
)

/*
  'bpmn:Association',
  'bpmn:BusinessRuleTask',
  'bpmn:DataInputAssociation',
  'bpmn:DataOutputAssociation',
  'bpmn:DataObjectReference',
  'bpmn:DataStoreReference',
  'bpmn:EndEvent',
  'bpmn:EventBasedGateway',
  'bpmn:ExclusiveGateway',
  'bpmn:IntermediateCatchEvent',
  'bpmn:ManualTask',
  'bpmn:ParallelGateway',
  'bpmn:Process',
  'bpmn:SequenceFlow',
  'bpmn:StartEvent',
  'bpmn:SubProcess',
  'bpmn:Task',
  'bpmn:TextAnnotation',
  'bpmn:UserTask'
*/

const (
	TYPE_GATEWAY  = "gateway"
	TYPE_EVENT    = "event"
	TYPE_TASK     = "task"
	TYPE_ACTIVITY = "activity"
	TYPE_NONE     = ""

	TAG_ROOT                    = "bpmn:definitions"
	TAG_COLABORATION            = "bpmn:collaboration"
	TAG_PROCESS                 = "bpmn:process"
	TAG_SUB_PROCESS             = "bpmn:subProcess"
	TAG_START_EVENT             = "bpmn:startEvent"
	TAG_END_EVENT               = "bpmn:endEvent"
	TAG_OUTGOING                = "bpmn:outgoing"
	TAG_INCOMING                = "bpmn:incoming"
	TAG_SEQUENCE_FLOW           = "bpmn:sequenceFlow"
	TAG_TASK                    = "bpmn:task"
	TAG_EVENT_BASED_GATEWAY     = "bpmn:eventBasedGateway"
	TAG_EVENT_EXCLUSIVE_GATEWAY = "bpmn:exclusiveGateway"
	TAG_EVENT_PARALLEL_GATEWAY  = "bpmn:parallelGateway"
	TAG_EVENT_COMPLEX_GATWAY    = "bpmn:complexGateway"
	TAG_DATA_INPUT_ASSOCIATION  = "bpmn:dataInputAssociation"
	TAG_DATA_OUTPUT_ASSOCIATION = "bpmn:dataOutputAssociation"
	TAG_DATA_OBJECT_REFERENCE   = "bpmn:dataObjectReference"

	ATTR_ID         = "id"
	ATTR_SOURCE_REF = "sourceRef"
	ATTR_TARGET_REF = "targetRef"
	ATTR_NAME       = "name"
)

/* Verifica si un elemento es una estructura gateway */
func isGateway(elem *etree.Element) (bool) {
	return strings.HasSuffix(elem.Tag, "Gateway")
}

/* Verifica si un elemento es un evento */
func isEvent(elem *etree.Element) (bool) {
	return strings.HasSuffix(elem.Tag, "Event")
}

/* Verifica si un elemento es una actividad */
func isActivity(elem *etree.Element) (bool) {
	activities := []string{"subProcess", "transaction", "task"}
	for _, v := range activities {
		if v == elem.Tag {
			return true
		}
	}
	return false
}

/* Obtiene el tipo de elemento */
func getTypeElement(elem *etree.Element) (string) {
	if isGateway(elem) {
		return TYPE_GATEWAY
	}
	if isEvent(elem) {
		return TYPE_EVENT
	}
	if isActivity(elem) {
		return TYPE_ACTIVITY
	}

	return TYPE_NONE
}

/* Funcion que crea la estructura de datos para el diagrama bpmn */

func NewDiagram(pathFilename string) DiagramBpmnIO {
	doc := etree.NewDocument()

	if err := doc.ReadFromFile(pathFilename); err != nil {
		panic(err)
	}

	//Se obtienen todas las sequencias del proceso
	sq := doc.SelectElement(TAG_ROOT).SelectElement(TAG_PROCESS).SelectElements(TAG_SEQUENCE_FLOW)

	return DiagramBpmnIO{documentXML: doc, sequences: sq}
}

/*
	Estructura para el diagrama BPMN de la herramienta bpmn.io
 */

type DiagramBpmnIO struct {
	documentXML *etree.Document
	sequences   []*etree.Element
}

/*
	Esta funcion es la encargada de obtener el elemento principal del diagrama XML
	es decir la etiqueta TAG_ROOT

	Retorna un puntero Element
 */
func (this DiagramBpmnIO) getRootElement() (*etree.Element) {
	return this.documentXML.SelectElement(TAG_ROOT)
}

/*
	Esta funcion es la encargada de obtener el elemento process donde estan ubicados todas
	las actividades del diagrama, es hija del elemento TAG_ROOT > TAG_PROCESS

	Retorna un puntero Element
 */
func (this DiagramBpmnIO) getProcessElement() (*etree.Element) {
	return this.getRootElement().SelectElement(TAG_PROCESS)
}

/*
	Esta funcion es usada para buscar cualquier elemento dentro de la etiquta TAG_PROCESS
	que coincida con el parametro id

	Retorna un puntero Element
 */
func (this DiagramBpmnIO) getElementByID(id string) (*etree.Element) {
	return this.getProcessElement().FindElement(`//[@id='` + id + `']`)
}

/*
	Esta funcion es usada para verificar si un elemento posee datos de entrada

	Retorna booleano
 */
func (this DiagramBpmnIO) HasDataInput(elem *etree.Element) (bool) {
	return len(elem.SelectElements(TAG_DATA_INPUT_ASSOCIATION)) > 0
}

/*
	Esta funcion obtiene todos los datos de entrada de un elemento

	Retorna slice de puntero element
 */
func (this DiagramBpmnIO) GetDataInputElement(elem *etree.Element) ([]*etree.Element) {
	data := elem.SelectElements(TAG_DATA_INPUT_ASSOCIATION)
	slice_data := make([]*etree.Element, 0)

	for _, input_asoc := range data {
		ref := input_asoc.SelectElement(ATTR_SOURCE_REF)
		if ref != nil {
			id_object_ref := ref.Text()
			slice_data = append(slice_data, this.getElementByID(id_object_ref))
		}
	}
	return slice_data
}

/*
	Esta funcion obtiene todoo el flujo del proceso, secuencial en un slice de elementos

	Retorna slice de puntero element
 */
func (this DiagramBpmnIO) GetSuccessionProcess() []*etree.Element {
	//ojo Sequence Element
	var sq_elem, elem_add *etree.Element
	s := make([]*etree.Element, 0)

	sq_elem = this.getBeginElement()

	elem_add = this.getElementByID(sq_elem.SelectAttrValue(ATTR_SOURCE_REF, ""))
	s = append(s, elem_add)

	for {
		if this.hasMoreElements(sq_elem) {
			sq_elem = this.getNextElement(sq_elem)

			elem_add = this.getElementByID(sq_elem.SelectAttrValue(ATTR_SOURCE_REF, ""))
			s = append(s, elem_add)

		} else {
			elem_add = this.getElementByID(sq_elem.SelectAttrValue(ATTR_TARGET_REF, ""))
			s = append(s, elem_add)
			break
		}
	}
	return s
}

/*
	Esta funcion obtiene el elemento principal basandose en TAG_START_EVENT

	Retorna un puntero element
 */
func (this DiagramBpmnIO) getBeginElement() (*etree.Element) {
	s_event := this.getProcessElement().SelectElement(TAG_START_EVENT)
	for _, v := range this.sequences {
		if v.SelectAttr(ATTR_SOURCE_REF).Value == s_event.SelectAttr(ATTR_ID).Value {
			return v
		}
	}
	return nil
}

/*
	Esta funcion obtiene el siguiente elemento en la sequencia del proceso

	Retorna un puntero element o nil
 */
func (this DiagramBpmnIO) getNextElement(previus *etree.Element) (*etree.Element) {
	for _, seq := range this.sequences {
		if seq.SelectAttr(ATTR_SOURCE_REF).Value == previus.SelectAttr(ATTR_TARGET_REF).Value {
			return seq
		}
	}
	return nil
}

/*
	Esta funcion verifica si existen mas elementos en la sequencia del proceso

	Retorna booleano
 */
func (this DiagramBpmnIO) hasMoreElements(previus *etree.Element) bool {
	return this.getNextElement(previus) != nil
}

/*
	Esta funcion obtiene la sequencia del proceso

	Retorna Retorna slice de puntero element
 */
func (this DiagramBpmnIO) GetSequences() ([]*etree.Element) {
	return this.sequences
}
