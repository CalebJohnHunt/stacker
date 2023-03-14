package main

import (
	"strings"

	sw "github.com/CalebJohnHunt/stacker"

	tea "github.com/charmbracelet/bubbletea"
)

type displayTextModel struct {
	text string
}

func (d *displayTextModel) Init() tea.Cmd { return nil }

func (d *displayTextModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "enter" {
		return d, sw.AddScene(&displayTextModel{text: d.text + string('a'+((d.text[len(d.text)-1]+1-'a')%26))})
	}
	if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "esc" {
		return d, sw.PopSceneSilent()
	}
	return d, nil
}

func (d *displayTextModel) View() string { return d.text }

type lister struct {
	list []tea.Model
	idx  int
}

func (l *lister) Init() tea.Cmd { return nil }
func (l *lister) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return l, tea.Quit
		case "up":
			l.idx--
		case "down":
			l.idx++
		}
	}
	var cmd tea.Cmd
	l.list[l.idx], cmd = l.list[l.idx].Update(msg)
	return l, cmd
}
func (l *lister) View() string {
	var sb strings.Builder
	for i, item := range l.list {
		if i == l.idx {
			sb.WriteByte('>')
		} else {
			sb.WriteByte(' ')
		}
		sb.WriteString(item.View())
		sb.WriteByte('\n')
	}
	return sb.String()
}

func main() {
	l := lister{
		list: []tea.Model{sw.NewSwitcher(&displayTextModel{"hello"}), sw.NewSwitcher(&displayTextModel{"world"})}}
	tea.NewProgram(&l).Run()
}
