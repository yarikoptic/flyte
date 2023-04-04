// Code generated by mockery v1.0.1. DO NOT EDIT.

package mocks

import (
	context "context"

	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	executors "github.com/flyteorg/flytepropeller/pkg/controller/executors"

	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/flyteorg/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
)

// Node is an autogenerated mock type for the Node type
type Node struct {
	mock.Mock
}

type Node_AbortHandler struct {
	*mock.Call
}

func (_m Node_AbortHandler) Return(_a0 error) *Node_AbortHandler {
	return &Node_AbortHandler{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnAbortHandler(ctx context.Context, execContext executors.ExecutionContext, dag executors.DAGStructure, nl executors.NodeLookup, currentNode v1alpha1.ExecutableNode, reason string) *Node_AbortHandler {
	c_call := _m.On("AbortHandler", ctx, execContext, dag, nl, currentNode, reason)
	return &Node_AbortHandler{Call: c_call}
}

func (_m *Node) OnAbortHandlerMatch(matchers ...interface{}) *Node_AbortHandler {
	c_call := _m.On("AbortHandler", matchers...)
	return &Node_AbortHandler{Call: c_call}
}

// AbortHandler provides a mock function with given fields: ctx, execContext, dag, nl, currentNode, reason
func (_m *Node) AbortHandler(ctx context.Context, execContext executors.ExecutionContext, dag executors.DAGStructure, nl executors.NodeLookup, currentNode v1alpha1.ExecutableNode, reason string) error {
	ret := _m.Called(ctx, execContext, dag, nl, currentNode, reason)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, executors.ExecutionContext, executors.DAGStructure, executors.NodeLookup, v1alpha1.ExecutableNode, string) error); ok {
		r0 = rf(ctx, execContext, dag, nl, currentNode, reason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Node_FinalizeHandler struct {
	*mock.Call
}

func (_m Node_FinalizeHandler) Return(_a0 error) *Node_FinalizeHandler {
	return &Node_FinalizeHandler{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnFinalizeHandler(ctx context.Context, execContext executors.ExecutionContext, dag executors.DAGStructure, nl executors.NodeLookup, currentNode v1alpha1.ExecutableNode) *Node_FinalizeHandler {
	c_call := _m.On("FinalizeHandler", ctx, execContext, dag, nl, currentNode)
	return &Node_FinalizeHandler{Call: c_call}
}

func (_m *Node) OnFinalizeHandlerMatch(matchers ...interface{}) *Node_FinalizeHandler {
	c_call := _m.On("FinalizeHandler", matchers...)
	return &Node_FinalizeHandler{Call: c_call}
}

// FinalizeHandler provides a mock function with given fields: ctx, execContext, dag, nl, currentNode
func (_m *Node) FinalizeHandler(ctx context.Context, execContext executors.ExecutionContext, dag executors.DAGStructure, nl executors.NodeLookup, currentNode v1alpha1.ExecutableNode) error {
	ret := _m.Called(ctx, execContext, dag, nl, currentNode)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, executors.ExecutionContext, executors.DAGStructure, executors.NodeLookup, v1alpha1.ExecutableNode) error); ok {
		r0 = rf(ctx, execContext, dag, nl, currentNode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Node_Initialize struct {
	*mock.Call
}

func (_m Node_Initialize) Return(_a0 error) *Node_Initialize {
	return &Node_Initialize{Call: _m.Call.Return(_a0)}
}

func (_m *Node) OnInitialize(ctx context.Context) *Node_Initialize {
	c_call := _m.On("Initialize", ctx)
	return &Node_Initialize{Call: c_call}
}

func (_m *Node) OnInitializeMatch(matchers ...interface{}) *Node_Initialize {
	c_call := _m.On("Initialize", matchers...)
	return &Node_Initialize{Call: c_call}
}

// Initialize provides a mock function with given fields: ctx
func (_m *Node) Initialize(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type Node_RecursiveNodeHandler struct {
	*mock.Call
}

func (_m Node_RecursiveNodeHandler) Return(_a0 executors.NodeStatus, _a1 error) *Node_RecursiveNodeHandler {
	return &Node_RecursiveNodeHandler{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Node) OnRecursiveNodeHandler(ctx context.Context, execContext executors.ExecutionContext, dag executors.DAGStructure, nl executors.NodeLookup, currentNode v1alpha1.ExecutableNode) *Node_RecursiveNodeHandler {
	c_call := _m.On("RecursiveNodeHandler", ctx, execContext, dag, nl, currentNode)
	return &Node_RecursiveNodeHandler{Call: c_call}
}

func (_m *Node) OnRecursiveNodeHandlerMatch(matchers ...interface{}) *Node_RecursiveNodeHandler {
	c_call := _m.On("RecursiveNodeHandler", matchers...)
	return &Node_RecursiveNodeHandler{Call: c_call}
}

// RecursiveNodeHandler provides a mock function with given fields: ctx, execContext, dag, nl, currentNode
func (_m *Node) RecursiveNodeHandler(ctx context.Context, execContext executors.ExecutionContext, dag executors.DAGStructure, nl executors.NodeLookup, currentNode v1alpha1.ExecutableNode) (executors.NodeStatus, error) {
	ret := _m.Called(ctx, execContext, dag, nl, currentNode)

	var r0 executors.NodeStatus
	if rf, ok := ret.Get(0).(func(context.Context, executors.ExecutionContext, executors.DAGStructure, executors.NodeLookup, v1alpha1.ExecutableNode) executors.NodeStatus); ok {
		r0 = rf(ctx, execContext, dag, nl, currentNode)
	} else {
		r0 = ret.Get(0).(executors.NodeStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, executors.ExecutionContext, executors.DAGStructure, executors.NodeLookup, v1alpha1.ExecutableNode) error); ok {
		r1 = rf(ctx, execContext, dag, nl, currentNode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type Node_SetInputsForStartNode struct {
	*mock.Call
}

func (_m Node_SetInputsForStartNode) Return(_a0 executors.NodeStatus, _a1 error) *Node_SetInputsForStartNode {
	return &Node_SetInputsForStartNode{Call: _m.Call.Return(_a0, _a1)}
}

func (_m *Node) OnSetInputsForStartNode(ctx context.Context, execContext executors.ExecutionContext, dag executors.DAGStructureWithStartNode, nl executors.NodeLookup, inputs *core.LiteralMap) *Node_SetInputsForStartNode {
	c_call := _m.On("SetInputsForStartNode", ctx, execContext, dag, nl, inputs)
	return &Node_SetInputsForStartNode{Call: c_call}
}

func (_m *Node) OnSetInputsForStartNodeMatch(matchers ...interface{}) *Node_SetInputsForStartNode {
	c_call := _m.On("SetInputsForStartNode", matchers...)
	return &Node_SetInputsForStartNode{Call: c_call}
}

// SetInputsForStartNode provides a mock function with given fields: ctx, execContext, dag, nl, inputs
func (_m *Node) SetInputsForStartNode(ctx context.Context, execContext executors.ExecutionContext, dag executors.DAGStructureWithStartNode, nl executors.NodeLookup, inputs *core.LiteralMap) (executors.NodeStatus, error) {
	ret := _m.Called(ctx, execContext, dag, nl, inputs)

	var r0 executors.NodeStatus
	if rf, ok := ret.Get(0).(func(context.Context, executors.ExecutionContext, executors.DAGStructureWithStartNode, executors.NodeLookup, *core.LiteralMap) executors.NodeStatus); ok {
		r0 = rf(ctx, execContext, dag, nl, inputs)
	} else {
		r0 = ret.Get(0).(executors.NodeStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, executors.ExecutionContext, executors.DAGStructureWithStartNode, executors.NodeLookup, *core.LiteralMap) error); ok {
		r1 = rf(ctx, execContext, dag, nl, inputs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
