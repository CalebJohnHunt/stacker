package main

import (
	"strings"

	sw "github.com/CalebJohnHunt/stacker"

	tea "github.com/charmbracelet/bubbletea"
)

type model1 struct {
	name string
	rd   string
}

func (m *model1) Init() tea.Cmd {
	return nil
}

func (m *model1) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "p":
			return m, sw.AddScene(&model1{name: m.name + "0"})
		case "i":
			return m, sw.AddScene(&inputLine{})
		case "I":
			return m, sw.AddSceneNoCallback(&inputLine{}) // This shouldn't return the model
		case "ctrl+p":
			return m, sw.PopScene()
		}
	case *model1:
		m.rd = msg.name
	case *inputLine:
		m.rd = string(msg.input)
		// default:
		// 	fmt.Printf("\n\n%T\n\n\n", msg)
	}
	return m, nil
}

func (m *model1) View() string {
	sb := strings.Builder{}
	sb.WriteString("My name is ")
	sb.WriteString(m.name)
	sb.WriteByte('\n')
	if m.rd != "" {
		sb.WriteString("Got returned data: ")
		sb.WriteString(m.rd)
	}
	return sb.String()
}
