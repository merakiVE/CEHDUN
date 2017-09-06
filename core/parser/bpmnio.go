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
	TAG_MESSAGE_FLOW            = "bpmn:messageFlow"
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

/* Funcion que crea la estructura de datos para el diagrama bpmn */

func NewDiagram() DiagramBpmnIO {
	doc := etree.NewDocument()
	return DiagramBpmnIO{documentXML: doc, sequences: make([]*etree.Element, 0), flows: make([]*etree.Element, 0)}
}

/*
	Estructura para el diagrama BPMN de la herramienta bpmn.io
 */

type DiagramBpmnIO struct {
	documentXML *etree.Document
	sequences   []*etree.Element
	flows       []*etree.Element
}

/* Funcion que carga el diagrama XML en forma de pathfile */
func (this *DiagramBpmnIO) ReadFromFile(filename string) {
	if err := this.documentXML.ReadFromFile(filename); err != nil {
		panic(err)
	}
	this.loadSequence()
}

/* Funcion que carga el diagrama XML en forma de string */
func (this *DiagramBpmnIO) ReadFromString(data string) {
	if err := this.documentXML.ReadFromString(data); err != nil {
		panic(err)
	}
	this.loadSequence()
}

/* Funcion que carga el diagrama XML en forma de byte */
func (this *DiagramBpmnIO) ReadFromBytes(bytes []byte) {
	if err := this.documentXML.ReadFromBytes(bytes); err != nil {
		panic(err)
	}
	this.loadSequence()
}

/* Funcion que carga la sequencia en un slice de la estructura */
func (this *DiagramBpmnIO) loadSequence() {
	//Se obtienen todas las sequencias del proceso
	//this.sequences = this.getProcessElement().SelectElements(TAG_SEQUENCE_FLOW)
	this.loadFlows()
}

/* Funcion que carga todos los flows en un slice de la estructura */
func (this *DiagramBpmnIO) loadFlows() {
	//s := make([]*etree.Element,0)

	//for _, eprocess := range this.getProcess() {
	//s = append(s, eprocess.FindElements(`[`+TAG_SEQUENCE_FLOW+`]`)...)
	//}
	this.flows = append(this.flows, this.getRootElement().FindElements(`[`+TAG_SEQUENCE_FLOW+`]`)...)
	this.flows = append(this.flows, this.getRootElement().FindElements(`[`+TAG_MESSAGE_FLOW+`]`)...)
}

/* Verifica si un elemento es una estructura gateway */
func (this DiagramBpmnIO) isGateway(elem *etree.Element) (bool) {
	return strings.HasSuffix(elem.Tag, "Gateway")
}

/* Verifica si un elemento es un evento */
func (this DiagramBpmnIO) isEvent(elem *etree.Element) (bool) {
	return strings.HasSuffix(elem.Tag, "Event")
}

/* Verifica si un elemento es una actividad */
func (this DiagramBpmnIO) isActivity(elem *etree.Element) (bool) {
	activities := []string{"subProcess", "transaction", "task"}
	for _, v := range activities {
		if v == elem.Tag {
			return true
		}
	}
	return false
}

/* Obtiene el tipo de elemento */
func (this DiagramBpmnIO) GetTypeElement(elem *etree.Element) (string) {
	if this.isGateway(elem) {
		return TYPE_GATEWAY
	}
	if this.isEvent(elem) {
		return TYPE_EVENT
	}
	if this.isActivity(elem) {
		return TYPE_ACTIVITY
	}

	return TYPE_NONE
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

func (this DiagramBpmnIO) getProcess() ([]*etree.Element) {
	return this.getRootElement().SelectElements(TAG_PROCESS)
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
	Esta funcion es usada para buscar cualquier elemento dentro de la etiquta TAG_ROOT
	que coincida con el atributo proporcionado

	Retorna un puntero Element
 */
func (this DiagramBpmnIO) getElementByAttr(atrib string, val string) (*etree.Element) {
	return this.getRootElement().FindElement(`//[@` + atrib + `='` + val + `']`)
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
	Esta funcion obtiene el atributo nombre del elemento

	Retorna un string con el nombre
 */
func (this DiagramBpmnIO) GetAttribute(elem *etree.Element, key string) (string) {
	return elem.SelectAttrValue(key, "unknown")
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
	Esta funcion obtiene todoo carriles del diagrama

	Retorna slice de puntero element
 */
func (this DiagramBpmnIO) GetLanes() []*etree.Element {
	//return this.getProcessElement()
	return nil
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
