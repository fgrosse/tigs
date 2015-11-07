package main

import (
	"fmt"
	"strings"

	"github.com/fgrosse/gotility"
)

type endpointTree struct {
	endpoints []endpoint
	nodes     map[string]*endpointNode
	rootNodes gotility.StringSet
}

type endpointNode struct {
	*endpoint
	children []*endpointNode
}

func newEndpointTree(endpoints []endpoint) endpointTree {
	t := endpointTree{
		endpoints: endpoints,
		nodes:     map[string]*endpointNode{},
		rootNodes: gotility.StringSet{},
	}

	for i := range endpoints {
		n := endpoints[i].Name
		t.nodes[n] = &endpointNode{&endpoints[i], []*endpointNode{}}
		t.rootNodes.Set(n)
	}

	return t
}

func (t endpointTree) process() error {
	// build up the inheritance tree
	for name, o := range t.nodes {
		if o.Extends == "" {
			continue
		}

		parent, ok := t.nodes[o.Extends]
		if !ok {
			return fmt.Errorf("endpoint %q extends an unknown endpoint %q", name, o.Extends)
		}

		parent.children = append(parent.children, t.nodes[name])
		t.rootNodes.Delete(name)
	}

	// check for any inheritance cycles
	for i := range t.endpoints {
		name := t.endpoints[i].Name
		trace := gotility.StringSlice{name}
		if err := t.checkForCycles(t.nodes[name], trace); err != nil {
			return err
		}
	}

	// finally inherit attributes in the tree
	for name := range t.rootNodes {
		for _, child := range t.nodes[name].children {
			t.inheritAttributes(t.nodes[name], child)
		}
	}

	return nil
}

func (t endpointTree) checkForCycles(node *endpointNode, trace gotility.StringSlice) error {
	for _, child := range node.children {
		if trace.Contains(child.Name) {
			trace.Add(child.Name)
			trace.Reverse()
			return fmt.Errorf("Detected inheritance cycle: %s", strings.Join(trace, " -> "))
		}

		trace.Add(child.Name)
		if err := t.checkForCycles(child, trace); err != nil {
			return err
		}
		trace.Delete(child.Name)
	}

	return nil
}

func (t endpointTree) inheritAttributes(parent, child *endpointNode) {
	if parent.Method != "" && child.Method == "" {
		child.Method = parent.Method
	}

	if parent.URL != "" && child.URL == "" {
		child.URL = parent.URL
	}

	for _, parentParameter := range parent.Parameters {
		keyExists := false
		for _, existingParameter := range child.Parameters {
			if existingParameter.Name == parentParameter.Name {
				keyExists = true
			}
		}

		if keyExists {
			// don't overwrite parameters that already exist in the children
			continue
		}

		child.Parameters = append(child.Parameters, parentParameter)
	}

	for _, grandChild := range child.children {
		t.inheritAttributes(child, grandChild)
	}
}
