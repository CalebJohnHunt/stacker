package stacker

import tea "github.com/charmbracelet/bubbletea"

func NewStacker(m tea.Model) tea.Model {
	if m == nil {
		return nil
	}
	return &Stacker{
		stack: []tea.Model{m}}
}

type pushSceneMsg struct {
	model tea.Model
}

func AddScene(m tea.Model) tea.Cmd {
	return func() tea.Msg {
		return pushSceneMsg{m}
	}
}

type popSceneMsg struct {
	silent bool
}

func PopScene() tea.Cmd {
	return func() tea.Msg {
		return popSceneMsg{}
	}
}

func PopSceneSilent() tea.Cmd {
	return func() tea.Msg {
		return popSceneMsg{silent: true}
	}
}

type Stacker struct {
	stack []tea.Model
}

func (m *Stacker) Init() tea.Cmd {
	return m.stack[len(m.stack)-1].Init()
}

func (m *Stacker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case pushSceneMsg:
		m.stack = append(m.stack, msg.model)
		return m, msg.model.Init()
	case popSceneMsg:
		// TODO:
		// If this stack is on a stack, it should somehow only pop itself...
		// I don't think it's possible with the current API/structs/etc.
		if len(m.stack) == 1 {
			return m, tea.Quit
		}
		curScene := m.stack[len(m.stack)-1]
		m.stack = m.stack[:len(m.stack)-1]
		if msg.silent {
			return m, nil
		}
		return m, func() tea.Msg { return curScene }
	}
	var b tea.Cmd
	m.stack[len(m.stack)-1], b = m.stack[len(m.stack)-1].Update(msg)
	return m, b
}

func (m *Stacker) View() string {
	return m.stack[len(m.stack)-1].View()
}
