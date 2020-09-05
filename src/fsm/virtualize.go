package fsm

import (
	"bytes"
	"fmt"
	"sort"
)

type VirtualizeType int

const (
	VirtualizeGraphviz VirtualizeType = iota
	VirtualizeMermaid
)

type Visualizer interface {
	Visualize(machine FSM) ([]byte, error)
}

var visualizers = make(map[VirtualizeType]Visualizer)

func init() {
	visualizers[VirtualizeMermaid] = MermaidVisualizer{}
	visualizers[VirtualizeGraphviz] = GraphvizVisualizer{}
}

func Virtualize(fsm FSM, t VirtualizeType) ([]byte, error) {
	visualizer, ok := visualizers[t]
	if !ok {
		return nil, ErrorUnknownVirtualizeType{}
	}
	return visualizer.Visualize(fsm)
}

type MermaidVisualizer struct{}

func (m MermaidVisualizer) Visualize(machine FSM) ([]byte, error) {
	fsm := machine.(*fsm)

	sortedEKeys := make([]eKey, 0, len(fsm.transitions))

	for k := range fsm.transitions {
		sortedEKeys = append(sortedEKeys, k)
	}

	sort.Slice(sortedEKeys, func(i, j int) bool {
		return sortedEKeys[i].src < sortedEKeys[j].src
	})

	buffer := bytes.Buffer{}
	buffer.WriteString("graph TD\n")

	for _, key := range sortedEKeys {
		val := fsm.transitions[key]
		buffer.WriteString(fmt.Sprintf("	_%s_-->|_%s_|_%s_\n", key.src, key.event, val))
	}

	return buffer.Bytes(), nil
}

type GraphvizVisualizer struct{}

func (g GraphvizVisualizer) Visualize(machine FSM) ([]byte, error) {
	fsm := machine.(*fsm)

	sortedEKeys := make([]eKey, 0, len(fsm.transitions))

	for k := range fsm.transitions {
		sortedEKeys = append(sortedEKeys, k)
	}

	sort.Slice(sortedEKeys, func(i, j int) bool {
		return sortedEKeys[i].src < sortedEKeys[j].src
	})

	buffer := bytes.Buffer{}
	states := make(map[string]int)

	buffer.WriteString("digraph fsm {\n")

	for _, key := range sortedEKeys {
		val := fsm.transitions[key]
		if key.src == fsm.current {
			states[key.src]++
			states[val]++
			buffer.WriteString(fmt.Sprintf("\"_%s_\"->\"_%s+\"[label=\"_%s_\"];\n", key.src, val, key.event))
		}
	}

	buffer.WriteString("\n")

	sortedStateKeys := make([]string, 0, len(states))

	for key := range states {
		sortedStateKeys = append(sortedStateKeys, key)
	}

	sort.Strings(sortedStateKeys)

	for _, key := range sortedStateKeys {
		buffer.WriteString(fmt.Sprintf("\"_%s_\";\n", key))
	}

	buffer.WriteString("}\n")

	return buffer.Bytes(), nil
}
