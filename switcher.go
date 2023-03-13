package stacker

import tea "github.com/charmbracelet/bubbletea"

func NewSwitcher(m tea.Model) tea.Model {
	return &Stacker{
		stack: []tea.Model{m}}
}

type pushSceneMsg struct {
	model tea.Model
}

type popSceneMsg struct {
}

type Stacker struct {
	stack []tea.Model
}

func (s *Stacker) Peek() tea.Model {
	return s.stack[len(s.stack)-1]
}

func (m *Stacker) Init() tea.Cmd {
	return m.stack[len(m.stack)-1].Init()
}

func (m *Stacker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case pushSceneMsg:
		m.stack = append(m.stack, msg.model)
		return m, msg.model.Init()
	case popSceneMsg:
		curScene := m.stack[len(m.stack)-1]
		m.stack = m.stack[:len(m.stack)-1]
		return m, func() tea.Msg { return curScene }
	}
	var b tea.Cmd
	m.stack[len(m.stack)-1], b = m.stack[len(m.stack)-1].Update(msg)
	return m, b
}

func (m *Stacker) View() string {
	return m.stack[len(m.stack)-1].View()
}

func AddScene(m tea.Model) tea.Cmd {
	return func() tea.Msg {
		return pushSceneMsg{m}
	}
}

func AddSceneNoCallback(m tea.Model) tea.Cmd {
	panic("Not implemented. Perhaps we have a scene{tea.model, options}")
}

func PopScene() tea.Cmd {
	return func() tea.Msg {
		return popSceneMsg{}
	}
}
