package stacker

import tea "github.com/charmbracelet/bubbletea"

func NewSwitcher(m tea.Model) tea.Model {
	return &Stacker{
		stack: []Scene{{m, nil}}}
}

type pushSceneMsg struct {
	scene Scene
}

type popSceneMsg struct {
	data any
}

type Scene struct {
	model    tea.Model
	callback cb
}

type cb func(any) tea.Cmd

type Stacker struct {
	stack []Scene
}

func (s *Stacker) Peek() tea.Model {
	return s.stack[len(s.stack)-1].model
}

func (m *Stacker) Init() tea.Cmd {
	return m.stack[len(m.stack)-1].model.Init()
}

func (m *Stacker) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var batch tea.Cmd = nil
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	case pushSceneMsg:
		batch = tea.Batch(msg.scene.model.Init())
		m.stack = append(m.stack, Scene{msg.scene.model, msg.scene.callback})
	case popSceneMsg:
		curScene := m.stack[len(m.stack)-1]
		m.stack = m.stack[:len(m.stack)-1]
		if curScene.callback != nil {
			batch = tea.Batch(batch, curScene.callback(msg.data))
		}
	}
	var b tea.Cmd
	m.stack[len(m.stack)-1].model, b = m.stack[len(m.stack)-1].model.Update(msg)
	batch = tea.Batch(batch, b)
	return m, batch
}

func (m *Stacker) View() string {
	return m.stack[len(m.stack)-1].model.View()
}

func AddScene(m tea.Model, callback cb) tea.Cmd {
	return func() tea.Msg {
		return pushSceneMsg{Scene{m, callback}}
	}
}

func AddSceneNoCallback(m tea.Model) tea.Cmd {
	return func() tea.Msg {
		return pushSceneMsg{Scene{m, nil}}
	}
}

func AddSceneSimpleSetter[T, MsgT any](m tea.Model, setter func(*MsgT, T)) tea.Cmd {
	return func() tea.Msg {
		return pushSceneMsg{Scene{m, SimpleSetterCallback(setter)}}
	}
}

func PopScene(data any) tea.Cmd {
	return func() tea.Msg {
		return popSceneMsg{data}
	}
}

func SimpleSetterCallback[T, MsgT any](setter func(*MsgT, T)) cb {
	s := new(MsgT)
	return func(data any) tea.Cmd {
		if d, ok := data.(T); ok {
			setter(s, d)
		}
		return func() tea.Msg {
			return *s
		}
	}
}
