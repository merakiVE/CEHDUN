package parser

import (
	"testing"
	"fmt"
)

const TEMPLATE_XML = `
<?xml version="1.0" encoding="UTF-8"?>
<bpmn:definitions xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:bpmn="http://www.omg.org/spec/BPMN/20100524/MODEL" xmlns:bpmndi="http://www.omg.org/spec/BPMN/20100524/DI" xmlns:dc="http://www.omg.org/spec/DD/20100524/DC" xmlns:di="http://www.omg.org/spec/DD/20100524/DI" id="Definitions_1" targetNamespace="http://bpmn.io/schema/bpmn">
  <bpmn:process id="Process_1" isExecutable="false">
    <bpmn:startEvent id="StartEvent_1">
      <bpmn:outgoing>SequenceFlow_14unrf1</bpmn:outgoing>
    </bpmn:startEvent>
    <bpmn:task id="Task_0hjkno0" name="@task1">
      <bpmn:incoming>SequenceFlow_14unrf1</bpmn:incoming>
      <bpmn:outgoing>SequenceFlow_1xvb71c</bpmn:outgoing>
    </bpmn:task>
    <bpmn:sequenceFlow id="SequenceFlow_14unrf1" sourceRef="StartEvent_1" targetRef="Task_0hjkno0" />
    <bpmn:sequenceFlow id="SequenceFlow_1xvb71c" sourceRef="Task_0hjkno0" targetRef="Task_0kblk1g" />
    <bpmn:task id="Task_0kblk1g" name="@task2">
      <bpmn:incoming>SequenceFlow_1xvb71c</bpmn:incoming>
      <bpmn:outgoing>SequenceFlow_1pplxrl</bpmn:outgoing>
    </bpmn:task>
    <bpmn:sequenceFlow id="SequenceFlow_1pplxrl" sourceRef="Task_0kblk1g" targetRef="Task_18ysb81" />
    <bpmn:task id="Task_1p01p5g" name="@task4">
      <bpmn:incoming>SequenceFlow_13s26qh</bpmn:incoming>
      <bpmn:outgoing>SequenceFlow_11htm28</bpmn:outgoing>
      <bpmn:dataOutputAssociation id="DataOutputAssociation_1vomhaw">
        <bpmn:targetRef>DataObjectReference_0fcehkj</bpmn:targetRef>
      </bpmn:dataOutputAssociation>
    </bpmn:task>
    <bpmn:sequenceFlow id="SequenceFlow_13s26qh" sourceRef="Task_18ysb81" targetRef="Task_1p01p5g" />
    <bpmn:dataObjectReference id="DataObjectReference_1gi64xz" name="@data_task3" dataObjectRef="DataObject_0l66wou" />
    <bpmn:dataObject id="DataObject_0l66wou" />
    <bpmn:dataObjectReference id="DataObjectReference_0bdry9s" name="@data2_task3" dataObjectRef="DataObject_1380jib" />
    <bpmn:dataObject id="DataObject_1380jib" />
    <bpmn:dataObjectReference id="DataObjectReference_0fcehkj" dataObjectRef="DataObject_18o8nam" />
    <bpmn:dataObject id="DataObject_18o8nam" />
    <bpmn:subProcess id="Task_18ysb81" name="@task3">
      <bpmn:incoming>SequenceFlow_1pplxrl</bpmn:incoming>
      <bpmn:outgoing>SequenceFlow_13s26qh</bpmn:outgoing>
      <bpmn:property name="__targetRef_placeholder" />
      <bpmn:dataInputAssociation />
      <bpmn:dataInputAssociation />
      <bpmn:dataInputAssociation id="DataInputAssociation_01ocnxq">
        <bpmn:sourceRef>DataObjectReference_1gi64xz</bpmn:sourceRef>
      </bpmn:dataInputAssociation>
      <bpmn:dataInputAssociation id="DataInputAssociation_0vlsm4d">
        <bpmn:sourceRef>DataObjectReference_0bdry9s</bpmn:sourceRef>
      </bpmn:dataInputAssociation>
    </bpmn:subProcess>
    <bpmn:sequenceFlow id="SequenceFlow_11htm28" sourceRef="Task_1p01p5g" targetRef="IntermediateThrowEvent_1i3hi10" />
    <bpmn:intermediateThrowEvent id="IntermediateThrowEvent_1i3hi10">
      <bpmn:incoming>SequenceFlow_11htm28</bpmn:incoming>
      <bpmn:outgoing>SequenceFlow_1150765</bpmn:outgoing>
      <bpmn:escalationEventDefinition />
    </bpmn:intermediateThrowEvent>
    <bpmn:endEvent id="EndEvent_0ob3mpg">
      <bpmn:incoming>SequenceFlow_1150765</bpmn:incoming>
    </bpmn:endEvent>
    <bpmn:sequenceFlow id="SequenceFlow_1150765" sourceRef="IntermediateThrowEvent_1i3hi10" targetRef="EndEvent_0ob3mpg" />
  </bpmn:process>
</bpmn:definitions>
`

func TestBpmnIO(t *testing.T) {

	d := NewDiagram()

	d.ReadFromString(TEMPLATE_XML)

	for _, n := range d.GetSuccessionProcess() {

		fmt.Println(n.Tag, d.GetTypeElement(n), d.HasDataInput(n), d.GetDataInputElement(n))
	}
}
